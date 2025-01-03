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

package schedulingtypes

import (
	"reflect"
	"strings"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtimeclient "sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	Pod = GetResourceKind(&corev1.Pod{})
)

var PodResource = &metav1.APIResource{
	Name:       GetPluralName(Pod),
	Group:      corev1.SchemeGroupVersion.Group,
	Version:    corev1.SchemeGroupVersion.Version,
	Kind:       Pod,
	Namespaced: true,
}

func GetResourceKind(obj runtimeclient.Object) string {
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Ptr {
		panic("All types must be pointers to structs.")
	}

	t = t.Elem()
	return t.Name()
}

func GetPluralName(name string) string {
	return strings.ToLower(name) + "s"
}
