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

package region

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	serviceName = "Service"
	regionId    = "region-id-1"
)

var (
	endpoint  = fmt.Sprintf("https://%s.%s.myhuaweicloud.com", strings.ToLower(serviceName), regionId)
	regionStr = `
SERVICE:
  - id: region-id-1
    endpoint: 'https://service.region-id-1.myhuaweicloud.com'
`
)

func EnvProviderTest(t *testing.T) {
	// test new
	p := NewEnvProvider(serviceName)
	assert.Equal(t, &EnvProvider{serviceName: strings.ToUpper(serviceName)}, p)
	// test not found
	region := p.GetRegion(regionId)
	assert.Nil(t, region)
	// test found
	err := os.Setenv(fmt.Sprintf("HUAWEICLOUD_SDK_REGION_%s_%s", serviceName, regionId), endpoint)
	assert.Nil(t, err)
	region = p.GetRegion(regionId)
	assert.Equal(t, NewRegion(regionId, endpoint), region)
}

func ProfileProviderTest(t *testing.T) {
	// test new
	p := NewProfileProvider(serviceName)
	assert.Equal(t, &ProfileProvider{serviceName: strings.ToUpper(serviceName)}, p)
	// get path
	dir, err := os.UserHomeDir()
	assert.Nil(t, err)
	filename := "test_regions.yaml"
	path := filepath.Join(dir, filename)
	err = os.Setenv("HUAWEICLOUD_SDK_REGIONS_FILE", path)
	assert.Nil(t, err)
	// create profile
	file, err := os.Create(path)
	assert.Nil(t, err)
	file.WriteString(regionStr)
	file.Close()
	defer os.Remove(path)
	// test found
	region := p.GetRegion(regionId)
	assert.Equal(t, region, NewRegion(regionId, endpoint))
}

func ProviderChainTest(t *testing.T) {
	// test new
	p := DefaultProviderChain(serviceName)
	p2 := &ProviderChain{
		serviceName: serviceName,
		providers: []IRegionProvider{
			&EnvProvider{serviceName: strings.ToUpper(serviceName)},
			&ProfileProvider{serviceName: strings.ToUpper(serviceName)},
		},
	}
	assert.Equal(t, p2, p)
	// test not found
	region := p.GetRegion("not-exist-1")
	assert.Nil(t, region)
	// test found
	region = p.GetRegion(regionId)
	assert.Equal(t, NewRegion(regionId, endpoint), region)

	region = DefaultProviderChain("NotExist").GetRegion(regionId)
	assert.Nil(t, region)
}

func TestRegionProvider(t *testing.T) {
	t.Run("EnvProviderTest", EnvProviderTest)
	t.Run("ProfileRegionTest", ProfileProviderTest)
	t.Run("ProviderChainTest", ProviderChainTest)
}
