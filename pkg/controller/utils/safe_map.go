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
	"sync"
)

type SafeMap struct {
	sync.RWMutex
	m map[string]interface{}
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		m: make(map[string]interface{}),
	}
}

func (s *SafeMap) Store(key string, value interface{}) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = value
}

func (s *SafeMap) Get(key string) (interface{}, bool) {
	s.RLock()
	defer s.RUnlock()
	value, ok := s.m[key]
	return value, ok
}

func (s *SafeMap) GetAll() []interface{} {
	s.RLock()
	defer s.RUnlock()
	var vals []interface{}
	for _, val := range s.m {
		vals = append(vals, val)
	}
	return vals
}

func (s *SafeMap) Delete(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.m, key)
}

func (s *SafeMap) DeleteAll() {
	s.Lock()
	defer s.Unlock()
	for key := range s.m {
		delete(s.m, key)
	}
}

func (s *SafeMap) Size() int {
	s.Lock()
	defer s.Unlock()
	return len(s.m)
}
