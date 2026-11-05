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
	assert.NoError(t, err)

	p := NewProfileProvider("Service")
	assert.Equal(t, "SERVICE", p.serviceName)
}

func TestProfileProvider_GetRegion(t *testing.T) {
	err := setRegionsFileEnv()
	assert.NoError(t, err)

	p := NewProfileProvider("Service1")
	region := p.GetRegion("region-id-1")
	assert.NotNil(t, region)
	assert.Equal(t, "region-id-1", region.Id)
	assert.Equal(t, []string{"https://service1.region-id-1.com"}, region.Endpoints)
}

func TestProfileProvider_GetRegion2(t *testing.T) {
	err := setRegionsFileEnv()
	assert.NoError(t, err)

	p := NewProfileProvider("Service2")
	region := p.GetRegion("region-id-2")
	assert.NotNil(t, region)
	assert.Equal(t, "region-id-2", region.Id)
	assert.Equal(t, []string{"https://service2.region-id-2.com", "https://service2.region-id-2.cn"}, region.Endpoints)
}

func TestProfileProvider_DoubleQuotedString(t *testing.T) {
	err := setRegionsFileEnv()
	assert.NoError(t, err)

	p := NewProfileProvider("Service3")
	region := p.GetRegion("region-id-3")
	assert.NotNil(t, region)
	assert.Equal(t, "region-id-3", region.Id)
	assert.Equal(t, []string{"https://service3.region-id-3.com:8080"}, region.Endpoints)
}

func TestProfileProvider_MultipleRegionsPerService(t *testing.T) {
	err := setRegionsFileEnv()
	assert.NoError(t, err)

	p := NewProfileProvider("Service3")
	region := p.GetRegion("region-id-4")
	assert.NotNil(t, region)
	assert.Equal(t, "region-id-4", region.Id)
	assert.Equal(t, []string{"https://service3.region-id-4.com", "https://service3.region-id-4.cn"}, region.Endpoints)
}

func TestProfileProvider_UrlWithQueryAndPort(t *testing.T) {
	err := setRegionsFileEnv()
	assert.NoError(t, err)

	p := NewProfileProvider("Service4")
	region := p.GetRegion("region-id-5")
	assert.NotNil(t, region)
	assert.Equal(t, "region-id-5", region.Id)
	assert.Equal(t, []string{"https://service4.region-id-5.com/path?q=1"}, region.Endpoints)
}

func TestProfileProvider_MixedEndpointStyles(t *testing.T) {
	err := setRegionsFileEnv()
	assert.NoError(t, err)

	p := NewProfileProvider("Service5")
	region6 := p.GetRegion("region-id-6")
	assert.NotNil(t, region6)
	assert.Equal(t, []string{"https://service5.region-id-6.com"}, region6.Endpoints)

	region7 := p.GetRegion("region-id-7")
	assert.NotNil(t, region7)
	assert.Equal(t, []string{"https://service5.region-id-7.com"}, region7.Endpoints)
}

func setRegionsFileEnv() error {
	absPath, err := os.Getwd()
	if err != nil {
		return err
	}
	file := path.Join(absPath, "test_regions.yaml")
	err = os.Setenv("HUAWEICLOUD_SDK_REGIONS_FILE", file)
	if err != nil {
		return err
	}
	return nil
}
