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
	"path"
	"testing"
)

func TestNewProfileProvider(t *testing.T) {
	err := setRegionsFileEnv()
	assert.Nil(t, err)

	p := NewProfileProvider("Service")
	assert.Equal(t, "SERVICE", p.serviceName)
}

func TestProfileProvider_GetRegion(t *testing.T) {
	err := setRegionsFileEnv()
	assert.Nil(t, err)

	p := NewProfileProvider("Service1")
	region := p.GetRegion("region-id-1")
	assert.NotNil(t, region)
	assert.Equal(t, "region-id-1", region.Id)
	assert.Equal(t, []string{"https://service1.region-id-1.com"}, region.Endpoints)
}

func TestProfileProvider_GetRegion2(t *testing.T) {
	err := setRegionsFileEnv()
	assert.Nil(t, err)

	p := NewProfileProvider("Service2")
	region := p.GetRegion("region-id-2")
	assert.NotNil(t, region)
	assert.Equal(t, "region-id-2", region.Id)
	assert.Equal(t, []string{"https://service2.region-id-2.com", "https://service2.region-id-2.cn"}, region.Endpoints)
}

func setRegionsFileEnv() error {
	absPath, err := os.Getwd()
	if err != nil {
		return err
	}
	file := path.Join(absPath, "regions.yaml")
	err = os.Setenv("HUAWEICLOUD_SDK_REGIONS_FILE", file)
	if err != nil {
		return err
	}
	return nil
}
