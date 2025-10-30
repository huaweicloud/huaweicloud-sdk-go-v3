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

package utils

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

type Case struct {
	Name           string
	X, Y, Expected int32
	Msg            string
}

func TestMax32(t *testing.T) {
	cases := []Case{
		{"test max32 with x > y", 5, 3, 5, "Expected max to be 5"},
		{"test max32 with y > x", 3, 5, 5, "Expected max to be 5"},
		{"test max32 with x == y", 5, 5, 5, "Expected max to be 5"},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := Max32(c.X, c.Y)
			assert.Equal(t, c.Expected, result, c.Msg)
		})
	}
}

func TestMin32(t *testing.T) {
	cases := []Case{
		{"test min32 with x < y", 3, 5, 3, "Expected min to be 3"},
		{"test min32 with y < x", 5, 3, 3, "Expected min to be 3"},
		{"test min32 with x == y", 5, 5, 5, "Expected min to be 5"},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := Min32(c.X, c.Y)
			assert.Equal(t, c.Expected, result, c.Msg)
		})
	}
}

func TestRandInt32(t *testing.T) {
	t.Run("test RandInt32 with valid range", func(t *testing.T) {
		min := int32(1)
		max := int32(10)
		result := RandInt32(min, max)
		assert.True(t, result >= min && result < max, "Expected result to be in [1, 10)")
	})

	t.Run("test RandInt32 with min > max", func(t *testing.T) {
		min := int32(10)
		max := int32(1)
		result := RandInt32(min, max)
		assert.Equal(t, max, result, "Expected result to be max when min > max")
	})

	t.Run("test RandInt32 with min <= 0", func(t *testing.T) {
		min := int32(-1)
		max := int32(10)
		result := RandInt32(min, max)
		assert.Equal(t, max, result, "Expected result to be max when min <= 0")
	})
}

func TestPow32(t *testing.T) {
	cases := []Case{
		{"test Pow32 with y = 0", 5, 0, 1, "Expected result to be 1 when y = 0"},
		{"test Pow32 with y = 3", 2, 3, 8, "Expected result to be 8"},
		{"test Pow32 with y = 2", 3, 2, 9, "Expected result to be 9"},
		{"test Pow32 with large exponent", 2, 31, int32(math.Exp2(31)), "Expected result to be 2^31"},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			result := Pow32(c.X, c.Y)
			assert.Equal(t, c.Expected, result, c.Msg)
		})
	}
}
