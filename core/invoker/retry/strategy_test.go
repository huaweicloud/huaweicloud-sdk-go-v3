// Copyright 2023 Huawei Technologies Co.,Ltd.
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

package retry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	runTimes         = 10
	baseDelayForTest = int32(100)
)

func TestDecorRelatedJitter(t *testing.T) {
	drj := NewDecorRelatedJitter()
	drj.SetBaseDelay(baseDelayForTest)

	runMultipleTimes(runTimes, func() {
		delay := drj.ComputeDelayBeforeNextRetry(10)
		assert.True(t, delay >= baseDelayForTest)
		assert.True(t, delay < baseDelayForTest*3)
	})
}

func TestEqualJitter(t *testing.T) {
	ej := NewEqualJitter()
	ej.SetBaseDelay(baseDelayForTest)

	runMultipleTimes(runTimes, func() {
		delay := ej.ComputeDelayBeforeNextRetry(3)
		assert.True(t, delay > baseDelayForTest*4)
		assert.True(t, delay < baseDelayForTest*8)
	})
}

func TestExponential(t *testing.T) {
	ex := NewExponential()
	ex.SetBaseDelay(baseDelayForTest)

	delay := ex.ComputeDelayBeforeNextRetry(3)
	assert.True(t, delay == baseDelayForTest*8)
}

func TestNone(t *testing.T) {
	n := None{}
	n.SetBaseDelay(baseDelayForTest)
	assert.Equal(t, int32(0), n.ComputeDelayBeforeNextRetry(1e9))
}

func TestRandomJitter(t *testing.T) {
	rj := NewRandomJitter()
	rj.SetBaseDelay(baseDelayForTest)

	runMultipleTimes(runTimes, func() {
		delay := rj.ComputeDelayBeforeNextRetry(3)
		assert.True(t, delay > 1)
		assert.True(t, delay < baseDelayForTest*8)
	})
}

func runMultipleTimes(times int, fn func()) {
	for i := 0; i < times; i++ {
		fn()
	}
}
