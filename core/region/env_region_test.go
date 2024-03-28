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
//go:build !linux
// +build !linux

package region

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewEnvProvider(t *testing.T) {
	p := NewEnvProvider("Service")
	assert.Equal(t, &EnvProvider{serviceName: "SERVICE"}, p)
}

func TestEnvProvider_GetRegion(t *testing.T) {
	p := NewEnvProvider("Service")
	reg := p.GetRegion("not-exist-1")
	assert.Nil(t, reg)
}

func TestEnvProvider_GetRegion2(t *testing.T) {
	p := NewEnvProvider("Service1")
	err := os.Setenv("HUAWEICLOUD_SDK_REGION_SERVICE1_REGION_ID_1", "https://service1.region-id-1.com")
	assert.Nil(t, err)
	reg := p.GetRegion("region-id-1")
	assert.NotNil(t, reg)
	assert.Equal(t, "region-id-1", reg.Id)
	assert.Equal(t, []string{"https://service1.region-id-1.com"}, reg.Endpoints)
}

func TestEnvProvider_GetRegion3(t *testing.T) {
	p := NewEnvProvider("Service2")
	err := os.Setenv("HUAWEICLOUD_SDK_REGION_SERVICE2_REGION_ID_2", "https://service2.region-id-2.com,https://service2.region-id-2.cn")
	assert.Nil(t, err)
	reg := p.GetRegion("region-id-2")
	assert.NotNil(t, reg)
	assert.Equal(t, "region-id-2", reg.Id)
	assert.Equal(t, []string{"https://service2.region-id-2.com", "https://service2.region-id-2.cn"}, reg.Endpoints)
}
