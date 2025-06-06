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

package provider

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMetadataCredentialProvider(t *testing.T) {
	p := NewMetadataCredentialProvider("test")
	assert.Equal(t, &MetadataCredentialProvider{credentialType: "test"}, p)
	_, err := p.GetCredentials()
	assert.ErrorContains(t, err, "unsupported credential type")
}

func TestBasicCredentialMetadataProvider(t *testing.T) {
	p := BasicCredentialMetadataProvider()
	assert.Equal(t, &MetadataBasicCredentialProvider{}, p)
}

func TestGlobalCredentialMetadataProvider(t *testing.T) {
	p := GlobalCredentialMetadataProvider()
	assert.Equal(t, &MetadataGlobalCredentialProvider{}, p)
}
