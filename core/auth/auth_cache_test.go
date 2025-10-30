// Copyright 2025 Huawei Technologies Co.,Ltd.
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

package auth

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestPutAndGet(t *testing.T) {
	c := getCache()

	t.Run("basic read", func(t *testing.T) {
		key := "test_key"
		value := "test_value"

		c.put(key, value)
		retrieved, ok := c.get(key)

		assert.True(t, ok, "The existing keys should be found")
		assert.Equal(t, value, retrieved, "The value read should be consistent with the value written")
	})

	t.Run("Read non-existent key", func(t *testing.T) {
		_, ok := c.get("non_existent_key")
		assert.False(t, ok, "Non-existent keys should return false")
	})
}

func TestConcurrentAccess(t *testing.T) {
	c := getCache()
	key := "concurrent_key"
	value := "race_test_value"

	var wg sync.WaitGroup
	wg.Add(2)

	// Concurrent write test
	t.Run("Concurrent write safety", func(t *testing.T) {
		go func() {
			defer wg.Done()
			c.put(key, value+"1")
		}()

		go func() {
			defer wg.Done()
			c.put(key, value+"2")
		}()

		wg.Wait()

		// The final value should be the value written last
		result, _ := c.get(key)
		assert.Contains(t, []string{value + "1", value + "2"}, result, "The result should be one of the two")
	})

	// Mixed reading and writing test
	t.Run("Mixed reading and writing", func(t *testing.T) {
		const workers = 100
		wg.Add(workers * 2)

		for i := 0; i < workers; i++ {
			go func(n int) {
				defer wg.Done()
				c.put(key, value+string(rune(n)))
			}(i)

			go func(n int) {
				defer wg.Done()
				_, _ = c.get(key)
			}(i)
		}

		wg.Wait()
		assert.NotPanics(t, func() { c.get(key) }, "Concurrent read and write should not cause panic")
	})
}
