// Copyright 2025 Huawei Technologies Co.,Ltd.
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

package auth

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestBaseCredentials_NeedUpdate(t *testing.T) {
	credentials, err := NewBaseCredentialsBuilder().SafeBuild()
	assert.NoError(t, err)
	// Manually specifying a security token
	credentials, err = NewBaseCredentialsBuilder().
		WithAk("ak").
		WithSk("sk").
		WithSecurityToken("token").
		SafeBuild()
	assert.NoError(t, err)
	assert.False(t, credentials.needUpdateSecurityTokenFromMetadata())
	// Automatically update the expired security token
	credentials.expiredAt = 1
	assert.True(t, credentials.needUpdateSecurityTokenFromMetadata())
	// The security token has not expired
	credentials.expiredAt = time.Now().Unix() + 10000
	assert.False(t, credentials.needUpdateSecurityTokenFromMetadata())
}

func TestBaseCredentials_NeedUpdateAuthToken(t *testing.T) {
	credentials, err := NewBaseCredentialsBuilder().SafeBuild()
	assert.NoError(t, err)
	// Without idp_id or id_token
	assert.False(t, credentials.needUpdateFederalAuthToken())
	credentials.IdpId = "idp_id"
	credentials.IdTokenFile = "file"
	// Get the auth token for the first time
	assert.True(t, credentials.needUpdateFederalAuthToken())
	// Automatically update the expired auth token
	credentials.SecurityToken = "sec-token"
	assert.True(t, credentials.needUpdateFederalAuthToken())
	// The auth token has not expired
	credentials.expiredAt = time.Now().Unix() + 10000
	assert.False(t, credentials.needUpdateFederalAuthToken())
}

func TestBaseCredentialsBuilder_SafeBuild(t *testing.T) {
	credentials, err := NewBaseCredentialsBuilder().
		WithAk("ak").
		WithSk("sk").
		WithIdpId("idp-id").
		WithIamEndpointOverride("endpoint").
		WithIdTokenFile("file").
		WithSecurityToken("xxx").SafeBuild()
	assert.NoError(t, err)
	assert.Equal(t, "ak", credentials.AK)
	assert.Equal(t, "sk", credentials.SK)
	assert.Equal(t, "idp-id", credentials.IdpId)
	assert.Equal(t, "endpoint", credentials.IamEndpoint)
	assert.Equal(t, "file", credentials.IdTokenFile)
	assert.Equal(t, "xxx", credentials.SecurityToken)
}

func TestBaseCredentials_updateAuthTokenByIdToken(t *testing.T) {
	var count int32
	mockServer := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-IAM-Trace-Id", "trace-id")
		w.Header().Set("X-Request-Id", "request-id")
		if atomic.LoadInt32(&count) == 0 {
			// v3.0/OS-AUTH/id-token/tokens
			assert.Equal(t, "idp_id", r.Header.Get("X-Idp-Id"))
			bytes, err := ioutil.ReadAll(r.Body)
			assert.NoError(t, err)
			body := strings.TrimSpace(string(bytes))
			assert.Equal(t, `{"auth":{"id_token":{"id":"id_token"}}}`, body)

			w.Header().Set("X-Subject-Token", "auth-token")
			_, err = w.Write([]byte(`{"token":{"expires_at":"2018-03-13T03:00:01.168000Z","methods":["mapped"]}}`))
			assert.NoError(t, err)
		} else {
			// v3.0/OS-CREDENTIAL/securitytokens
			bytes, err := ioutil.ReadAll(r.Body)
			assert.NoError(t, err)
			body := strings.TrimSpace(string(bytes))
			assert.Equal(t, `{"auth":{"identity":{"methods":["token"],"token":{"id":"auth-token","duration_seconds":21600}}}}`, body)

			_, err = w.Write([]byte(`{"credential":{"access":"ak","expires_at":"2020-01-08T03:50:07.574000Z","secret":"sk","securitytoken":"sec-token"}}`))
			assert.NoError(t, err)
		}

		atomic.AddInt32(&count, 1)
	}))
	defer mockServer.Close()

	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "BaseTestIdTokenFile")
	err := os.WriteFile(filePath, []byte("id_token"), 0600)
	filePath, err = filepath.Abs(filePath)
	assert.NoError(t, err)

	cred, err := NewBaseCredentialsBuilder().
		WithIdpId("idp_id").
		WithIdTokenFile(filePath).
		WithIamEndpointOverride(mockServer.URL).
		SafeBuild()
	assert.NoError(t, err)
	client := impl.NewDefaultHttpClient(config.DefaultHttpConfig().WithIgnoreSSLVerification(true))
	err = cred.updateAuthTokenByIdToken(client)
	assert.NoError(t, err)
	assert.Equal(t, "ak", cred.AK)
	assert.Equal(t, "sk", cred.SK)
	assert.Equal(t, "sec-token", cred.SecurityToken)
}

func TestBaseCredentials_IsDerivedAuth(t *testing.T) {
	credentials, err := NewBaseCredentialsBuilder().WithDerivedPredicate(func(request *request.DefaultHttpRequest) bool {
		return true
	}).SafeBuild()
	assert.NoError(t, err)
	assert.True(t, credentials.IsDerivedAuth(nil))
}

func TestBaseCredentials_needUpdateSecurityTokenFromMetadata(t *testing.T) {
	credentials, err := NewBaseCredentialsBuilder().SafeBuild()
	assert.NoError(t, err)
	assert.True(t, credentials.needUpdateSecurityTokenFromMetadata())
}

func TestBaseCredentials_getIdToken(t *testing.T) {
	tmpDir := t.TempDir()
	filePath := filepath.Join(tmpDir, "BaseTestIdTokenFile")
	err := os.WriteFile(filePath, []byte(""), 0600)
	filePath, err = filepath.Abs(filePath)
	assert.NoError(t, err)

	credentials, err := NewBaseCredentialsBuilder().WithIdpId("id").WithIdTokenFile(filePath).SafeBuild()
	assert.NoError(t, err)
	_, err = credentials.getIdToken()
	assert.Errorf(t, err, "id token is empty")
}

func TestBaseCredentials_baseProcessAuthRequest(t *testing.T) {
	credentials, err := NewBaseCredentialsBuilder().
		WithAk("ak").
		WithSk("sk").
		WithDerivedPredicate(func(httpRequest *request.DefaultHttpRequest) bool {
			return true
		}).SafeBuild()
	assert.NoError(t, err)
	credentials.derivedAuthServiceName = "service"
	credentials.regionId = "globe"
	builder := request.NewHttpRequestBuilder()
	err = credentials.baseProcessAuthRequest(builder, builder.Build())
	assert.NoError(t, err)
}

func TestBaseCredentials_SafeBuild(t *testing.T) {
	_, err := NewBaseCredentialsBuilder().WithIdpId("id").SafeBuild()
	assert.Errorf(t, err, "IdTokenFile is required when using IdpId&IdTokenFile")

	_, err = NewBaseCredentialsBuilder().WithIdTokenFile("file").SafeBuild()
	assert.Errorf(t, err, "IdpId is required when using IdpId&IdTokenFile")
}

// Deprecated
func TestBaseCredentials_Build(t *testing.T) {
	assert.Panicsf(t, func() {
		NewBaseCredentialsBuilder().WithIdpId("id").Build()
	}, "IdTokenFile is required when using IdpId&IdTokenFile")
}
