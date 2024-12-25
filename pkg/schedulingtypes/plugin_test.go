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
	"testing"

	"sigs.k8s.io/kubefed/pkg/controller/utils"
)

func TestUpdateOverridesMap(t *testing.T) {
	cluster := "cluster1"

	testCases := map[string]struct {
		overridesMap utils.OverridesMap
		replicasMap  map[string]int64
		expected     map[string]int64
	}{
		"Retain other overrides when removing replica override for unscheduled": {
			overridesMap: utils.OverridesMap{
				cluster: utils.ClusterOverrides{
					{
						Path:  replicasPath,
						Value: int64(0),
					},
					{
						Path:  "/ultimate/answer",
						Value: int64(42),
					},
				},
			},
			replicasMap: make(map[string]int64),
			expected: map[string]int64{
				"/ultimate/answer": 42,
			},
		},
		"Update existing replica override": {
			overridesMap: utils.OverridesMap{
				cluster: utils.ClusterOverrides{
					{
						Path:  replicasPath,
						Value: int64(0),
					},
					{
						Path:  "/ultimate/answer",
						Value: int64(42),
					},
				},
			},
			replicasMap: map[string]int64{
				cluster: 5,
			},
			expected: map[string]int64{
				"/ultimate/answer": 42,
				replicasPath:       5,
			},
		},
		"Add new replica override": {
			overridesMap: utils.OverridesMap{
				cluster: utils.ClusterOverrides{
					{
						Path:  "/ultimate/answer",
						Value: int64(42),
					},
				},
			},
			replicasMap: map[string]int64{
				cluster: 0,
			},
			expected: map[string]int64{
				"/ultimate/answer": 42,
				replicasPath:       0,
			},
		},
	}

	for testName, tc := range testCases {
		t.Run(testName, func(t *testing.T) {
			updateOverridesMap(tc.overridesMap, tc.replicasMap)
			actual := make(map[string]int64)
			for _, override := range tc.overridesMap[cluster] {
				actual[override.Path] = override.Value.(int64)
			}
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Fatalf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}
