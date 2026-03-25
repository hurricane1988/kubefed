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

package utils

import (
	"context"
	"time"

	"github.com/pkg/errors"
	pkgruntime "k8s.io/apimachinery/pkg/runtime"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"

	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/client/apiutil"

	"sigs.k8s.io/kubefed/pkg/client/generic/scheme"
)

func NewGenericInformer(config *rest.Config, namespace string, obj runtimeclient.Object, resyncPeriod time.Duration, triggerFunc func(runtimeclient.Object)) (cache.Store, cache.Controller, error) {
	return NewGenericInformerWithEventHandler(config, namespace, obj, resyncPeriod, NewTriggerOnAllChanges(triggerFunc))
}

func NewGenericInformerWithEventHandler(config *rest.Config, namespace string, obj runtimeclient.Object, resyncPeriod time.Duration, resourceEventHandlerFuncs *cache.ResourceEventHandlerFuncs) (cache.Store, cache.Controller, error) {
	// Extract GroupVersionKind (GVK) from the provided runtime object.
	gvk, err := apiutil.GVKForObject(obj, scheme.Scheme)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get GVK for object")
	}

	// Initialize an HTTP Client based on the REST config.
	httpClient, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create HTTP client")
	}

	// Build a RESTMapper to resolve GVK to a RESTMapping.
	mapper, err := apiutil.NewDynamicRESTMapper(config, httpClient)
	if err != nil {
		return nil, nil, errors.Wrap(err, "could not create RESTMapper")
	}

	mapping, err := mapper.RESTMapping(gvk.GroupKind(), gvk.Version)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to get REST mapping")
	}

	// Create a GVK-specific REST Client.
	// The 6 arguments are: gvk, forceDisableProtoBuf, isUnstructured, baseConfig, codecs, httpClient.
	restClient, err := apiutil.RESTClientForGVK(
		gvk,
		false, // forceDisableProtoBuf: use Protobuf if available.
		false, // isUnstructured: set to false as we are using typed objects.
		config,
		scheme.Codecs,
		httpClient,
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to create REST client")
	}

	// Instantiate the List object type required for unmarshaling API responses.
	listGVK := gvk.GroupVersion().WithKind(gvk.Kind + "List")
	listObj, err := scheme.Scheme.New(listGVK)
	if err != nil {
		return nil, nil, errors.Wrapf(err, "failed to instantiate list object for %s", listGVK)
	}

	// Configure InformerOptions with a context-aware ListerWatcher.
	options := cache.InformerOptions{
		// ListWithContextFunc is the modern replacement for ListFunc.
		ListerWatcher: &cache.ListWatch{
			ListWithContextFunc: func(ctx context.Context, opts metav1.ListOptions) (pkgruntime.Object, error) {
				results := listObj.DeepCopyObject()
				// Only apply namespace filtering if the resource is namespaced and a namespace is provided.
				isNamespaceScoped := namespace != "" && mapping.Scope.Name() != meta.RESTScopeNameRoot

				err := restClient.Get().
					NamespaceIfScoped(namespace, isNamespaceScoped).
					Resource(mapping.Resource.Resource).
					VersionedParams(&opts, scheme.ParameterCodec).
					Do(ctx).
					Into(results)
				return results, err
			},
			// WatchWithContextFunc is the modern replacement for WatchFunc.
			WatchFuncWithContext: func(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
				opts.Watch = true
				isNamespaceScoped := namespace != "" && mapping.Scope.Name() != meta.RESTScopeNameRoot
				return restClient.Get().
					NamespaceIfScoped(namespace, isNamespaceScoped).
					Resource(mapping.Resource.Resource).
					VersionedParams(&opts, scheme.ParameterCodec).
					Watch(ctx)
			},
		},
		ObjectType:   obj, // The type of object the informer will cache.
		Handler:      resourceEventHandlerFuncs,
		ResyncPeriod: resyncPeriod,
	}

	// Create the Informer using the modern Options pattern.
	store, controller := cache.NewInformerWithOptions(options)
	return store, controller, nil
}
