// Copyright 2022 Huawei Technologies Co.,Ltd.
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

package region

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestNewEnvProvider(t *testing.T) {
	p := NewEnvProvider(serviceName)
	assert.Equal(t, &EnvProvider{serviceName: strings.ToUpper(serviceName)}, p)
}

func TestEnvProvider_GetRegion(t *testing.T) {
	p := NewEnvProvider(serviceName)
	reg := p.GetRegion(regionId)
	assert.Nil(t, reg)
}

func TestEnvProvider_GetRegion2(t *testing.T) {
	p := NewEnvProvider(serviceName)
	err := os.Setenv(fmt.Sprintf("HUAWEICLOUD_SDK_REGION_%s_%s",
		serviceName, strings.ToUpper(strings.Replace(regionId, "-", "_", -1))), endpoint)
	assert.Nil(t, err)
	reg := p.GetRegion(regionId)
	assert.Equal(t, NewRegion(regionId, endpoint), reg)
}
