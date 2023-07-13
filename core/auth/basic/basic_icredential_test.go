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

package basic

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCredentials_NeedUpdate(t *testing.T) {
	credentials := NewCredentialsBuilder().Build()
	// Manually specifying a security token
	credentials = NewCredentialsBuilder().WithAk("ak").WithSk("sk").WithSecurityToken("token").Build()
	assert.False(t, credentials.needUpdateSecurityToken())
	// Automatically update the expired security token
	credentials.expiredAt = 1
	assert.True(t, credentials.needUpdateSecurityToken())
	// The security token has not expired
	credentials.expiredAt = time.Now().Unix() + 120
	assert.False(t, credentials.needUpdateSecurityToken())
}

func TestCredentials_NeedUpdateAuthToken(t *testing.T) {
	credentials := NewCredentialsBuilder().Build()
	// Without idp_id or id_token
	assert.False(t, credentials.needUpdateAuthToken())
	credentials.IdpId = "idp_id"
	credentials.IdTokenFile = "file"
	// Get the auth token for the first time
	assert.True(t, credentials.needUpdateAuthToken())
	// Automatically update the expired auth token
	credentials.authToken = "auth_token"
	assert.True(t, credentials.needUpdateAuthToken())
	// The auth token has not expired
	credentials.expiredAt = time.Now().Unix() + 120
	assert.False(t, credentials.needUpdateAuthToken())
}

func TestCredentialsBuilder_Build(t *testing.T) {
	defer func() {
		err := recover()
		assert.IsType(t, err, &sdkerr.CredentialsTypeError{})
		assert.Equal(t, err.(*sdkerr.CredentialsTypeError).ErrorMessage, "ProjectId is required when using IdpId&IdTokenFile")
	}()

	NewCredentialsBuilder().WithIdpId("id").WithIdTokenFile("file").Build()
}

func TestCredentialsBuilder_Build1(t *testing.T) {
	defer func() {
		err := recover()
		assert.IsType(t, err, &sdkerr.CredentialsTypeError{})
		assert.Equal(t, err.(*sdkerr.CredentialsTypeError).ErrorMessage, "IdTokenFile is required when using IdpId&IdTokenFile")
	}()

	NewCredentialsBuilder().WithIdpId("id").Build()
}

func TestCredentialsBuilder_Build2(t *testing.T) {
	defer func() {
		err := recover()
		assert.IsType(t, err, &sdkerr.CredentialsTypeError{})
		assert.Equal(t, err.(*sdkerr.CredentialsTypeError).ErrorMessage, "IdpId is required when using IdpId&IdTokenFile")
	}()

	NewCredentialsBuilder().WithIdTokenFile("file").Build()
}

func TestCredentialsBuilder_Build3(t *testing.T) {
	NewCredentialsBuilder().WithIdpId("id").WithIdTokenFile("file").WithProjectId("projectId").Build()
}
