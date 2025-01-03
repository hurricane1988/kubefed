/*
Copyright 2024 The CodeFuture Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sync

import (
	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog/v2"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	corev1 "k8s.io/api/core/v1"

	"sigs.k8s.io/kubefed/pkg/apis/core/typeconfig"
	genericclient "sigs.k8s.io/kubefed/pkg/client/generic"
	"sigs.k8s.io/kubefed/pkg/controller/sync/version"
	"sigs.k8s.io/kubefed/pkg/controller/utils"
)

// FederatedResourceAccessor provides a way to retrieve and visit
// logical federated resources (e.g. FederatedConfigMap)
type FederatedResourceAccessor interface {
	Run(stopChan <-chan struct{})
	HasSynced() bool
	FederatedResource(qualifiedName utils.QualifiedName) (federatedResource FederatedResource, possibleOrphan bool, err error)
	VisitFederatedResources(visitFunc func(obj interface{}))
}

type resourceAccessor struct {
	limitedScope      bool
	typeConfig        typeconfig.Interface
	targetIsNamespace bool
	fedNamespace      string

	// The informer for the federated type.
	federatedStore      cache.Store
	federatedController cache.Controller

	// The informer used to source namespaces for templates of
	// federated namespaces.  Will only be initialized if
	// targetIsNamespace=true.
	namespaceStore      cache.Store
	namespaceController cache.Controller

	fedNamespaceAPIResource *metav1.APIResource

	// The informer used to source federated namespaces used in
	// determining placement for namespaced resources.  Will only be
	// initialized if the target resource is namespaced.
	fedNamespaceStore      cache.Store
	fedNamespaceController cache.Controller

	// Manages propagated versions
	versionManager *version.Manager

	// Records events on the federated resource
	eventRecorder record.EventRecorder
	// ctx is the context that governs the Manager's operations, allowing for graceful shutdowns or cancellations.
	ctx context.Context
	// immediate indicates whether the Manager should propagate version information immediately or wait for a synchronization cycle.
	immediate bool
}

func NewFederatedResourceAccessor(ctx context.Context, immediate bool, controllerConfig *utils.ControllerConfig, typeConfig typeconfig.Interface, fedNamespaceAPIResource *metav1.APIResource, client genericclient.Client, enqueueObj func(runtimeclient.Object), eventRecorder record.EventRecorder) (FederatedResourceAccessor, error) {
	a := &resourceAccessor{
		limitedScope:            controllerConfig.LimitedScope(),
		typeConfig:              typeConfig,
		targetIsNamespace:       typeConfig.GetTargetType().Kind == utils.NamespaceKind,
		fedNamespace:            controllerConfig.KubeFedNamespace,
		fedNamespaceAPIResource: fedNamespaceAPIResource,
		eventRecorder:           eventRecorder,
	}

	targetNamespace := controllerConfig.TargetNamespace

	federatedTypeAPIResource := typeConfig.GetFederatedType()
	federatedTypeClient, err := utils.NewResourceClient(controllerConfig.KubeConfig, &federatedTypeAPIResource)
	if err != nil {
		return nil, err
	}
	a.federatedStore, a.federatedController = utils.NewResourceInformer(federatedTypeClient, targetNamespace, &federatedTypeAPIResource, enqueueObj)

	if a.targetIsNamespace {
		// Initialize an informer for namespaces.  The namespace
		// containing a federated namespace resource is used as the
		// template for target resources in member clusters.
		namespaceAPIResource := typeConfig.GetTargetType()
		namespaceTypeClient, err := utils.NewResourceClient(controllerConfig.KubeConfig, &namespaceAPIResource)
		if err != nil {
			return nil, err
		}
		a.namespaceStore, a.namespaceController = utils.NewResourceInformer(namespaceTypeClient, targetNamespace, &namespaceAPIResource, enqueueObj)
	}

	if typeConfig.GetNamespaced() {
		fedNamespaceEnqueue := func(fedNamespaceObj runtimeclient.Object) {
			// When a federated namespace changes, every resource in
			// the namespace needs to be reconciled.
			//
			// TODO(marun) Consider optimizing this to only reconcile
			// contained resources in response to a change in
			// placement for the federated namespace.
			namespace := utils.NewQualifiedName(fedNamespaceObj).Namespace
			for _, rawObj := range a.federatedStore.List() {
				obj := rawObj.(runtimeclient.Object)
				qualifiedName := utils.NewQualifiedName(obj)
				if qualifiedName.Namespace == namespace {
					enqueueObj(obj)
				}
			}
		}
		// Initialize an informer for federated namespaces.  Placement
		// for a resource is computed as the intersection of resource
		// and federated namespace placement.
		fedNamespaceClient, err := utils.NewResourceClient(controllerConfig.KubeConfig, fedNamespaceAPIResource)
		if err != nil {
			return nil, err
		}
		a.fedNamespaceStore, a.fedNamespaceController = utils.NewResourceInformer(fedNamespaceClient, targetNamespace, fedNamespaceAPIResource, fedNamespaceEnqueue)
	}

	a.versionManager = version.NewVersionManager(ctx, immediate, client, typeConfig.GetFederatedNamespaced(), typeConfig.GetFederatedType().Kind, typeConfig.GetTargetType().Kind, targetNamespace)

	return a, nil
}

func (a *resourceAccessor) Run(stopChan <-chan struct{}) {
	go a.versionManager.Sync(stopChan)
	go a.federatedController.Run(stopChan)
	if a.namespaceController != nil {
		go a.namespaceController.Run(stopChan)
	}
	if a.fedNamespaceController != nil {
		go a.fedNamespaceController.Run(stopChan)
	}
}

func (a *resourceAccessor) HasSynced() bool {
	kind := a.typeConfig.GetFederatedType().Kind
	if !a.versionManager.HasSynced() {
		klog.V(2).Infof("Version manager for %s not synced", kind)
		return false
	}
	if !a.federatedController.HasSynced() {
		klog.V(2).Infof("Informer for %s not synced", kind)
		return false
	}
	if a.namespaceController != nil && !a.namespaceController.HasSynced() {
		klog.V(2).Infof("Namespace informer for %s not synced", kind)
		return false
	}
	if a.fedNamespaceController != nil && !a.fedNamespaceController.HasSynced() {
		klog.V(2).Infof("FederatedNamespace informer for %s not synced", kind)
		return false
	}
	return true
}

func (a *resourceAccessor) FederatedResource(eventSource utils.QualifiedName) (FederatedResource, bool, error) {
	if a.targetIsNamespace && a.isSystemNamespace(eventSource.Name) {
		klog.V(7).Infof("Ignoring system namespace %q", eventSource.Name)
		return nil, false, nil
	}

	kind := a.typeConfig.GetFederatedType().Kind

	// Most federated resources have the same name as their targets.
	targetName := utils.QualifiedName{
		Namespace: eventSource.Namespace,
		Name:      eventSource.Name,
	}
	federatedName := utils.QualifiedName{
		Namespace: utils.NamespaceForResource(eventSource.Namespace, a.fedNamespace),
		Name:      eventSource.Name,
	}

	// A federated type for namespace "foo" is namespaced
	// (e.g. "foo/foo"). An event sourced from a namespace in the host
	// or member clusters will have the name "foo", and an event
	// sourced from a federated resource will have the name "foo/foo".
	// In order to ensure object retrieval from the informers, it is
	// necessary to derive the target name and federated name from the
	// event source.
	if a.targetIsNamespace {
		eventSourceIsTarget := eventSource.Namespace == ""
		if eventSourceIsTarget {
			// Ensure the federated name is namespace qualified.
			federatedName.Namespace = federatedName.Name
		} else {
			// Ensure the target name is not namespace qualified.
			targetName.Namespace = ""
		}
	}

	key := federatedName.String()

	resource, err := utils.ObjFromCache(a.federatedStore, kind, key)
	if err != nil {
		return nil, false, err
	}
	if resource == nil {
		// If the target is a namespace and the event source has a
		// namespace, the event source is guaranteed to be a
		// FederatedNamespace.
		sourceIsFederatedNamespace := a.targetIsNamespace && eventSource.Namespace != ""

		// The lack of a federated resource indicates that the event
		// source may be an orphaned resource that still has the
		// managed label.
		possibleOrphan := !sourceIsFederatedNamespace
		return nil, possibleOrphan, nil
	}

	var namespace *unstructured.Unstructured
	if a.targetIsNamespace {
		if federatedName.Namespace != federatedName.Name {
			// A FederatedNamespace is only valid for propagation
			// if it has the same name as the containing namespace.
			a.eventRecorder.Eventf(
				resource, corev1.EventTypeWarning,
				"InvalidName", "The name of a federated namespace must match the name of its containing namespace.")
			return nil, false, nil
		}
		namespace, err = utils.ObjFromCache(a.namespaceStore, a.typeConfig.GetTargetType().Kind, targetName.String())
		if err != nil {
			return nil, false, err
		}
		if namespace == nil {
			// The namespace containing the FederatedNamespace was deleted.
			return nil, false, nil
		}
	}

	var fedNamespace *unstructured.Unstructured
	if a.typeConfig.GetNamespaced() {
		fedNamespaceName := utils.QualifiedName{Namespace: federatedName.Namespace, Name: federatedName.Namespace}
		fedNamespace, err = utils.ObjFromCache(a.fedNamespaceStore, a.fedNamespaceAPIResource.Kind, fedNamespaceName.String())
		if err != nil {
			return nil, false, err
		}
		// If fedNamespace is nil, the resources in member clusters
		// will be removed.
	}

	return &federatedResource{
		limitedScope:      a.limitedScope,
		typeConfig:        a.typeConfig,
		targetIsNamespace: a.targetIsNamespace,
		targetName:        targetName,
		federatedKind:     kind,
		federatedName:     federatedName,
		federatedResource: resource,
		versionManager:    a.versionManager,
		namespace:         namespace,
		fedNamespace:      fedNamespace,
		eventRecorder:     a.eventRecorder,
	}, false, nil
}

func (a *resourceAccessor) VisitFederatedResources(visitFunc func(obj interface{})) {
	for _, obj := range a.federatedStore.List() {
		visitFunc(obj)
	}
}

func (a *resourceAccessor) isSystemNamespace(namespace string) bool {
	// TODO(font): Need a configurable or discoverable list of namespaces
	// to not propagate beyond just the default system namespaces e.g.
	switch namespace {
	case "kube-system", "kube-public", "default", a.fedNamespace:
		return true
	default:
		return false
	}
}
