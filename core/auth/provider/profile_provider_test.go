// Copyright 2024 Huawei Technologies Co.,Ltd.
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
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

const (
	MockAk            = "MockAk"
	MockSk            = "MockSk"
	MockProjectId     = "MockProjectId"
	MockDomainId      = "MockDomainId"
	MockSecurityToken = "MockSecurityToken"
	MockIamEndpoint   = "MockIamEndpoint"
	MockIdpId         = "MockIdpId"
	MockIdTokenFile   = "MockIdTokenFile"
)

var credentialStr = `
[basic]
ak = MockAk
sk = MockSk
project_id = MockProjectId
security_token = MockSecurityToken
iam_endpoint = MockIamEndpoint    

[global]                          
ak = MockAk
sk = MockSk
idp_id = MockIdpId
id_token_file = MockIdTokenFile
domain_id = MockDomainId
security_token = MockSecurityToken
iam_endpoint = MockIamEndpoint

[error]
ak = ak
sk = sk
`

func TestNewProfileCredentialProvider(t *testing.T) {
	p := NewProfileCredentialProvider("test")
	assert.Equal(t, &ProfileCredentialProvider{credentialType: "test"}, p)
}

func TestBasicCredentialProfileProvider(t *testing.T) {
	p := BasicCredentialProfileProvider()
	assert.Equal(t, &ProfileCredentialProvider{credentialType: basicCredentialType}, p)
}

func TestGlobalCredentialProfileProvider(t *testing.T) {
	p := GlobalCredentialProfileProvider()
	assert.Equal(t, &ProfileCredentialProvider{credentialType: globalCredentialType}, p)
}

func TestProfileCredentialProvider_GetCredentials(t *testing.T) {
	dir, err := os.UserHomeDir()
	assert.Nil(t, err)
	filename := "test_credentials"
	path := filepath.Join(dir, ".huaweicloud", filename)
	err = os.Setenv("HUAWEICLOUD_SDK_CREDENTIALS_FILE", path)
	assert.Nil(t, err)

	file, err := os.Create(path)
	assert.Nil(t, err)
	err = file.Chmod(0600)
	assert.Nil(t, err)
	_, err = file.WriteString(credentialStr)
	assert.Nil(t, err)
	err = file.Close()
	assert.Nil(t, err)
	defer func(name string) {
		err = os.Remove(name)
		assert.Nil(t, err)
	}(path)

	basicProvider := BasicCredentialProfileProvider()
	basicCred, err := basicProvider.GetCredentials()
	assert.Nil(t, err)
	expectedBasicCred := basic.NewCredentialsBuilder().
		WithAk(MockAk).
		WithSk(MockSk).
		WithProjectId(MockProjectId).
		WithSecurityToken(MockSecurityToken).
		WithIamEndpointOverride(MockIamEndpoint).
		Build()
	assert.Equal(t, expectedBasicCred, basicCred)

	globalProvider := GlobalCredentialProfileProvider()
	globalCred, err := globalProvider.GetCredentials()
	assert.Nil(t, err)
	expectedGlobalCred := global.NewCredentialsBuilder().
		WithIdpId(MockIdpId).
		WithIdTokenFile(MockIdTokenFile).
		WithDomainId(MockDomainId).
		WithIamEndpointOverride(MockIamEndpoint).
		Build()
	assert.Equal(t, expectedGlobalCred, globalCred)

	neProvider := NewProfileCredentialProvider("notexist")
	neCred, err := neProvider.GetCredentials()
	assert.EqualError(t, err,
		"{\"ErrorMessage\": \"unsupported credential type: notexist\"}")
	assert.Nil(t, neCred)
}
