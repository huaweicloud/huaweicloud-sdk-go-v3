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

func TestNewProviderChain(t *testing.T) {
	chain := NewProviderChain(serviceName, []IRegionProvider{
		&EnvProvider{serviceName: strings.ToUpper(serviceName)},
		&ProfileProvider{serviceName: strings.ToUpper(serviceName)},
	})
	assert.Equal(t, &ProviderChain{
		serviceName: serviceName,
		providers: []IRegionProvider{
			&EnvProvider{serviceName: strings.ToUpper(serviceName)},
			&ProfileProvider{serviceName: strings.ToUpper(serviceName)},
		},
	}, chain)
}

func TestDefaultProviderChain(t *testing.T) {
	chain := DefaultProviderChain(serviceName)
	assert.Equal(t, &ProviderChain{
		serviceName: serviceName,
		providers: []IRegionProvider{
			&EnvProvider{serviceName: strings.ToUpper(serviceName)},
			&ProfileProvider{serviceName: strings.ToUpper(serviceName)},
		},
	}, chain)
}

func TestProviderChain_GetRegion(t *testing.T) {
	chain := DefaultProviderChain(serviceName)
	reg := chain.GetRegion("not-exist-1")
	assert.Nil(t, reg)
}

func TestProviderChain_GetRegion2(t *testing.T) {
	chain := DefaultProviderChain("NotExist")
	reg := chain.GetRegion(serviceName)
	assert.Nil(t, reg)
}

func TestProviderChain_GetRegion3(t *testing.T) {
	dir, err := os.UserHomeDir()
	assert.Nil(t, err)
	filename := "test_regions.yaml"
	path := filepath.Join(dir, filename)
	err = os.Setenv("HUAWEICLOUD_SDK_REGIONS_FILE", path)
	assert.Nil(t, err)

	file, err := os.Create(path)
	assert.Nil(t, err)
	_, err = file.WriteString(regionStr)
	assert.Nil(t, err)
	err = file.Close()
	assert.Nil(t, err)
	defer func(name string) {
		err = os.Remove(name)
		assert.Nil(t, err)
	}(path)

	chain := DefaultProviderChain(serviceName)
	reg := chain.GetRegion(regionId)
	assert.Equal(t, NewRegion(regionId, endpoint), reg)
}
