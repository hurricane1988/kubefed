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

package dispatch

import (
	"context"

	"github.com/pkg/errors"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/klog/v2"

	"sigs.k8s.io/kubefed/pkg/client/generic"
	"sigs.k8s.io/kubefed/pkg/controller/utils"
)

type isNamespaceInHostClusterFunc func(clusterObj runtimeclient.Object) bool

type CheckUnmanagedDispatcher interface {
	OperationDispatcher

	CheckRemovedOrUnlabeled(clusterName string, isHostNamespace isNamespaceInHostClusterFunc)
}

type checkUnmanagedDispatcherImpl struct {
	dispatcher *operationDispatcherImpl

	targetGVK  schema.GroupVersionKind
	targetName utils.QualifiedName
}

func NewCheckUnmanagedDispatcher(clientAccessor clientAccessorFunc, targetGVK schema.GroupVersionKind, targetName utils.QualifiedName) CheckUnmanagedDispatcher {
	dispatcher := newOperationDispatcher(clientAccessor, nil)
	return &checkUnmanagedDispatcherImpl{
		dispatcher: dispatcher,
		targetGVK:  targetGVK,
		targetName: targetName,
	}
}

func (d *checkUnmanagedDispatcherImpl) Wait() (bool, error) {
	return d.dispatcher.Wait()
}

// CheckRemovedOrUnlabeled checks that a resource either does not
// exist in the given cluster, or if it does exist, that it does not
// have the managed label.
func (d *checkUnmanagedDispatcherImpl) CheckRemovedOrUnlabeled(clusterName string, isHostNamespace isNamespaceInHostClusterFunc) {
	d.dispatcher.incrementOperationsInitiated()
	const op = "check for deletion of resource or removal of managed label from"
	const opContinuous = "Checking for deletion of resource or removal of managed label from"
	go d.dispatcher.clusterOperation(clusterName, op, func(client generic.Client) utils.ReconciliationStatus {
		targetName := d.targetNameForCluster(clusterName)

		klog.V(2).Infof(eventTemplate, opContinuous, d.targetGVK.Kind, targetName, clusterName)

		clusterObj := &unstructured.Unstructured{}
		clusterObj.SetGroupVersionKind(d.targetGVK)
		err := client.Get(context.Background(), clusterObj, targetName.Namespace, targetName.Name)
		if apierrors.IsNotFound(err) {
			return utils.StatusAllOK
		}
		if err != nil {
			wrappedErr := d.wrapOperationError(err, clusterName, op)
			runtime.HandleError(wrappedErr)
			return utils.StatusError
		}
		if clusterObj.GetDeletionTimestamp() != nil {
			if isHostNamespace(clusterObj) {
				return utils.StatusAllOK
			}
			err = errors.Errorf("resource is pending deletion")
			wrappedErr := d.wrapOperationError(err, clusterName, op)
			runtime.HandleError(wrappedErr)
			return utils.StatusError
		}
		if !utils.HasManagedLabel(clusterObj) {
			return utils.StatusAllOK
		}
		err = errors.Errorf("resource still has the managed label")
		wrappedErr := d.wrapOperationError(err, clusterName, op)
		runtime.HandleError(wrappedErr)
		return utils.StatusError
	})
}

func (d *checkUnmanagedDispatcherImpl) wrapOperationError(err error, clusterName, operation string) error {
	return wrapOperationError(err, operation, d.targetGVK.Kind, d.targetNameForCluster(clusterName).String(), clusterName)
}

func (d *checkUnmanagedDispatcherImpl) targetNameForCluster(clusterName string) utils.QualifiedName {
	return utils.QualifiedNameForCluster(clusterName, d.targetName)
}
