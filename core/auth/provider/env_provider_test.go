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
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewEnvCredentialProvider(t *testing.T) {
	p := NewEnvCredentialProvider("test")
	assert.Equal(t, &EnvCredentialProvider{credentialType: "test"}, p)
}

func TestBasicCredentialEnvProvider(t *testing.T) {
	p := BasicCredentialEnvProvider()
	assert.Equal(t, &EnvCredentialProvider{credentialType: basicCredentialType}, p)
}

func TestGlobalCredentialEnvProvider(t *testing.T) {
	p := GlobalCredentialEnvProvider()
	assert.Equal(t, &EnvCredentialProvider{credentialType: globalCredentialType}, p)
}

func TestEnvCredentialProvider_GetCredentials(t *testing.T) {
	// TestAkSkIsEmpty
	p := BasicCredentialEnvProvider()
	credentials, err := p.GetCredentials()
	assert.EqualError(t, err, "{\"ErrorMessage\": \"AK&SK or IdpId&IdTokenFile does not exist\"}")
	assert.Nil(t, credentials)
	err = resetEnv()
	assert.Nil(t, err)
}

func TestEnvCredentialProvider_GetCredentials2(t *testing.T) {
	// TestCredentialTypeIsEmpty
	p := NewEnvCredentialProvider("")
	credentials, err := p.GetCredentials()
	assert.EqualError(t, err, "{\"ErrorMessage\": \"credential type is empty\"}")
	assert.Nil(t, credentials)
	err = resetEnv()
	assert.Nil(t, err)
}

func TestEnvCredentialProvider_GetCredentials3(t *testing.T) {
	// TestBasicCredentials
	err := os.Setenv("HUAWEICLOUD_SDK_AK", MockAk)
	assert.Nil(t, err)
	err = os.Setenv("HUAWEICLOUD_SDK_SK", MockSk)
	assert.Nil(t, err)
	p := BasicCredentialEnvProvider()
	credentials, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, basic.NewCredentialsBuilder().WithAk(MockAk).WithSk(MockSk).Build(), credentials)

	p = BasicCredentialEnvProvider()
	assert.Nil(t, err)
	err = os.Setenv("HUAWEICLOUD_SDK_PROJECT_ID", MockProjectId)
	assert.Nil(t, err)
	credentials, err = p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, basic.NewCredentialsBuilder().WithAk(MockAk).WithSk(MockSk).WithProjectId(MockProjectId).Build(), credentials)
	err = resetEnv()
	assert.Nil(t, err)
}

func TestEnvCredentialProvider_GetCredentials4(t *testing.T) {
	// TestGlobalCredentials
	err := setAkSkEnv()
	assert.Nil(t, err)
	p := GlobalCredentialEnvProvider()
	credentials, err := p.GetCredentials()
	assert.Nil(t, err)
	assert.Equal(t, global.NewCredentialsBuilder().WithAk(MockAk).WithSk(MockSk).Build(), credentials)
	err = resetEnv()
	assert.Nil(t, err)
}

func TestEnvCredentialProvider_GetCredentials5(t *testing.T) {
	err := setAkSkEnv()
	assert.Nil(t, err)
	p := BasicCredentialEnvProvider()
	err = os.Setenv("HUAWEICLOUD_SDK_IDP_ID", MockIdpId)
	assert.Nil(t, err)
	err = os.Setenv("HUAWEICLOUD_SDK_ID_TOKEN_FILE", MockIdTokenFile)
	assert.Nil(t, err)
	err = os.Setenv("HUAWEICLOUD_SDK_PROJECT_ID", MockProjectId)
	assert.Nil(t, err)
	credentials, err := p.GetCredentials()
	assert.Nil(t, err)
	expected := basic.NewCredentialsBuilder().WithIdpId(MockIdpId).WithIdTokenFile(MockIdTokenFile).WithProjectId(MockProjectId).Build()
	assert.Equal(t, expected, credentials)
	err = resetEnv()
	assert.Nil(t, err)
}

func resetEnv() error {
	envs := []string{
		"HUAWEICLOUD_SDK_AK",
		"HUAWEICLOUD_SDK_SK",
		"HUAWEICLOUD_SDK_PROJECT_ID",
		"HUAWEICLOUD_SDK_DOMAIN_ID",
		"HUAWEICLOUD_SDK_CREDENTIAL_TYPE",
		"HUAWEICLOUD_SDK_IDP_ID",
		"HUAWEICLOUD_SDK_ID_TOKEN_FILE",
	}
	for _, env := range envs {
		err := os.Setenv(env, "")
		if err != nil {
			return err
		}
	}
	return nil
}

func setAkSkEnv() (err error) {
	err = os.Setenv("HUAWEICLOUD_SDK_AK", MockAk)
	err = os.Setenv("HUAWEICLOUD_SDK_SK", MockSk)
	return err
}
