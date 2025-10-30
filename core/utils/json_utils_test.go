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
	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshal(t *testing.T) {
	type S struct {
		Str string `json:"str"`
	}

	marshal, err := Marshal(S{Str: "1 + 2 > 2"})
	assert.NoError(t, err)
	assert.Equal(t, "{\"str\":\"1 + 2 > 2\"}\n", string(marshal))
}

func TestUnmarshal(t *testing.T) {
	type N struct {
		Num jsoniter.Number `json:"num"`
	}

	n := N{}
	err := Unmarshal([]byte(`{"num":10}`), &n)
	assert.NoError(t, err)
	assert.Equal(t, "10", n.Num.String())
}
