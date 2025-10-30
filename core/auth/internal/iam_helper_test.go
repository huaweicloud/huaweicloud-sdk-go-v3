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

package internal

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestIamEndpoint(t *testing.T) {
	endpoint := GetIamEndpoint()
	assert.Equal(t, DefaultIamEndpoint, endpoint)

	err := os.Setenv(IamEndpointEnv, "https://endpoint")
	assert.NoError(t, err)
	endpoint = GetIamEndpoint()
	assert.Equal(t, "https://endpoint", endpoint)

	err = os.Setenv(IamEndpointEnv, "endpoint")
	assert.NoError(t, err)
	endpoint = GetIamEndpoint()
	assert.Equal(t, "https://endpoint", endpoint)

	endpoint1 := GetIamEndpoint()
	endpoint2 := GetIamEndpointById("test")
	assert.Equal(t, endpoint1, endpoint2)
	endpoint3 := GetIamEndpointById("cn-north-4")
	assert.NotEmpty(t, endpoint3)
}
