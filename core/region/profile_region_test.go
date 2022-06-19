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
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

var regionStr = `
SERVICE:
  - id: region-id-1
    endpoint: 'https://service.region-id-1.myhuaweicloud.com'
`

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
	reg := p.GetRegion(regionId)
	assert.Equal(t, reg, NewRegion(regionId, endpoint))
}

func TestNewProfileProvider(t *testing.T) {
	p := NewProfileProvider(serviceName)
	assert.Equal(t, &ProfileProvider{serviceName: strings.ToUpper(serviceName)}, p)
}

func TestProfileProvider_GetRegion1(t *testing.T) {
	dir, err := os.UserHomeDir()
	assert.Nil(t, err)
	filename := "test_regions.yaml"
	path := filepath.Join(dir, filename)
	err = os.Setenv("HUAWEICLOUD_SDK_REGIONS_FILE", path)
	assert.Nil(t, err)

	file, err := os.Create(path)
	assert.Nil(t, err)
	file.WriteString(regionStr)
	file.Close()
	defer os.Remove(path)

	p1 := NewProfileProvider("NotExist")
	reg1 := p1.GetRegion(serviceName)
	assert.Nil(t, reg1)

	p2 := NewProfileProvider(serviceName)
	reg2 := p2.GetRegion("not-exist-3")
	assert.Nil(t, reg2)

	p3 := NewProfileProvider(serviceName)
	reg3 := p3.GetRegion(regionId)
	assert.Equal(t, NewRegion(regionId, endpoint), reg3)
}
