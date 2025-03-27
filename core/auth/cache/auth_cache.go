// Copyright 2020 Huawei Technologies Co.,Ltd.
//
// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cache

import (
	"sync"
)

type Cache interface {
	PutAuth(key, value string)
	GetAuth(key string) (string, bool)
}

var (
	instance Cache
	once     sync.Once
)

func GetCache() Cache {
	once.Do(func() {
		instance = &cacheImpl{
			data: make(map[string]string),
		}
	})
	return instance
}

type cacheImpl struct {
	data map[string]string
	mu   sync.RWMutex
}

func (s *cacheImpl) PutAuth(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *cacheImpl) GetAuth(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	val, ok := s.data[key]
	return val, ok
}
