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
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/signer/algorithm"
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
)

func TestGlobalCredentials_ProcessAuthParams(t *testing.T) {
	cases := []struct {
		Name, RegionId, Data, AK, SK, ExpectedDomainId string
		StatusCode                                     int
		ExpectedError                                  string
	}{
		{"Bad Request", "region-id-1", "{\"error_code\":\"XXX.001\",\"error_msg\":\"Bad Request\"}", "ak", "sk", "", 400,
			"failed to get domain id automatically, {\"status_code\":400,\"request_id\":\"\",\"error_code\":\"XXX.001\",\"error_message\":\"Bad Request, X-IAM-Trace-Id=trace-id\",\"encoded_authorization_message\":\"\",\"details\":null}"},
		{"One Domain", "region-id-1", "{\"domains\":[{\"id\":\"domain_id\"}]}", "ak", "sk", "domain_id", 200, ""},
		{"No Domain", "region-id-2", "{\"domains\":[]}", "ak2", "sk2", "", 200,
			"failed to get domain id automatically," +
				" X-IAM-Trace-Id=trace-id, please confirm that you have 'iam:users:getUser' permission," +
				" or set domain id manually:" +
				" global.NewCredentialsBuilder().WithAk(ak).WithSk(sk).WithDomainId(domainId).SafeBuild()"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("X-IAM-Trace-Id", "trace-id")
				w.WriteHeader(c.StatusCode)
				_, err := fmt.Fprintln(w, c.Data)
				assert.NoError(t, err)
			}))
			defer ts.Close()

			credentials, err := NewGlobalCredentialsBuilder().
				WithAk(c.AK).
				WithSk(c.SK).
				WithIamEndpointOverride(ts.URL).
				SafeBuild()
			assert.NoError(t, err)

			client := impl.NewDefaultHttpClient(config.DefaultHttpConfig().WithIgnoreSSLVerification(true))
			_, err = credentials.ProcessAuthParams(client, c.RegionId)
			if c.ExpectedError != "" {
				assert.ErrorContains(t, err, c.ExpectedError)
			} else {
				assert.Equal(t, c.ExpectedDomainId, credentials.DomainId)
			}
		})
	}
}

func TestGlobalCredentials_ProcessAuthParams2(t *testing.T) {
	credentials, err := NewGlobalCredentialsBuilder().WithAk("ak1").WithSk("sk1").SafeBuild()
	assert.NoError(t, err)
	client := impl.NewDefaultHttpClient(config.DefaultHttpConfig().WithSigningAlgorithm("test"))
	_, err = credentials.ProcessAuthParams(client, "region")
	assert.Error(t, err)

	credentials, err = NewGlobalCredentialsBuilder().WithAk("ak").WithSk("sk").SafeBuild()
	assert.NoError(t, err)
	credentials.DomainId = "domain-id"
	cred, err := credentials.ProcessAuthParams(nil, "")
	assert.NoError(t, err)
	credentials, ok := cred.(*GlobalCredentials)
	assert.True(t, ok)
	assert.Equal(t, "domain-id", credentials.DomainId)

	credentials, err = NewGlobalCredentialsBuilder().WithAk("ak").WithSk("sk").SafeBuild()
	assert.NoError(t, err)
	getCache().put("ak", "id")
	cred, err = credentials.ProcessAuthParams(nil, "region")
	assert.NoError(t, err)
	credentials, ok = cred.(*GlobalCredentials)
	assert.True(t, ok)
	assert.Equal(t, "id", credentials.DomainId)
}

func TestGlobalCredentials_ProcessAuthParams3(t *testing.T) {
	var status int
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("X-IAM-Trace-Id", "trace-id")

		if r.URL.Path == "/v3/auth/domains" {
			_, err := fmt.Fprintln(w, "{\"domains\":[]}")
			assert.NoError(t, err)
		} else {
			var err error
			w.Header().Set("x-request-id", "request-id")
			w.WriteHeader(status)
			if status == 200 {
				_, err = fmt.Fprintln(w, "{\"account_id\":\"AccountId\",\"principal_urn\":\"PrincipalUrn\",\"principal_id\":\"PrincipalId\"}")
			} else {
				_, err = fmt.Fprintln(w, "{\"error_code\":\"code\",\"error_msg\":\"Server error\"}")
			}
			assert.NoError(t, err)
		}
	}))
	defer ts.Close()
	err := os.Setenv("HUAWEICLOUD_SDK_STS_ENDPOINT", ts.URL)
	assert.NoError(t, err)
	err = os.Setenv("HUAWEICLOUD_SDK_IAM_ENDPOINT", ts.URL)
	assert.NoError(t, err)

	status = 200
	credentials, err := NewGlobalCredentialsBuilder().WithAk("ak3").WithSk("sk3").SafeBuild()
	assert.NoError(t, err)

	client := impl.NewDefaultHttpClient(config.DefaultHttpConfig().WithIgnoreSSLVerification(true))
	_, err = credentials.ProcessAuthParams(client, "region3")
	assert.NoError(t, err)
	assert.Equal(t, "AccountId", credentials.DomainId)

	status = 404
	credentials, err = NewGlobalCredentialsBuilder().WithAk("ak3_1").WithSk("sk3_1").SafeBuild()
	assert.NoError(t, err)
	_, err = credentials.ProcessAuthParams(client, "region3_1")
	assert.Errorf(t, err, "failed to get domain id from %s: 404, requestId: request-id", ts.URL)

	status = 500
	credentials, err = NewGlobalCredentialsBuilder().WithAk("ak3_2").WithSk("sk3_2").SafeBuild()
	assert.NoError(t, err)
	_, err = credentials.ProcessAuthParams(client, "region3_2")
	assert.ErrorContains(t, err, "failed to get domain id")
	assert.ErrorContains(t, err, "Server error")
}

func TestGlobalCredentials_FederalAuth(t *testing.T) {
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
	filePath := filepath.Join(tmpDir, "GlobalTestIdTokenFile")
	err := os.WriteFile(filePath, []byte("id_token"), 0600)
	filePath, err = filepath.Abs(filePath)
	assert.NoError(t, err)

	cred, err := NewGlobalCredentialsBuilder().
		WithIdpId("idp_id").
		WithIdTokenFile(filePath).
		WithIamEndpointOverride(mockServer.URL).
		SafeBuild()
	assert.NoError(t, err)
	client := impl.NewDefaultHttpClient(config.DefaultHttpConfig().WithIgnoreSSLVerification(true))
	req := request.NewHttpRequestBuilder().
		WithMethod("GET").
		WithPath("/test").
		WithEndpoint(mockServer.URL).
		WithSigningAlgorithm(algorithm.GetDefaultSigningAlgorithm()).
		Build()
	authReq, err := cred.ProcessAuthRequest(client, req)
	assert.NoError(t, err)
	assert.Equal(t, "sec-token", authReq.GetHeaderParams()["X-Security-Token"])
	assert.Equal(t, "ak", cred.AK)
	assert.Equal(t, "sk", cred.SK)
	assert.Equal(t, "sec-token", cred.SecurityToken)
}

func TestGlobalCredentials_ProcessDerivedAuthParams(t *testing.T) {
	credentials, err := NewGlobalCredentialsBuilder().WithAk("ak").WithSk("sk").SafeBuild()
	assert.NoError(t, err)
	credentials.ProcessDerivedAuthParams("name", "region")
	assert.Equal(t, "name", credentials.derivedAuthServiceName)
	assert.Equal(t, "globe", credentials.regionId)
}

func TestGlobalCredentials_ProcessAuthRequest(t *testing.T) {
	credentials, err := NewGlobalCredentialsBuilder().
		WithAk("ak").
		WithSk("sk").
		WithSecurityToken("xxx").
		WithDomainId("domain-id").
		SafeBuild()
	assert.NoError(t, err)
	req := request.NewHttpRequestBuilder().WithSigningAlgorithm(algorithm.GetDefaultSigningAlgorithm()).Build()
	req, err = credentials.ProcessAuthRequest(nil, req)
	assert.NoError(t, err)
	assert.Equal(t, req.GetHeaderParams()["X-Domain-Id"], "domain-id")
	assert.Contains(t, req.GetHeaderParams(), "Authorization")

	credentials, err = NewGlobalCredentialsBuilder().
		WithIdpId("idp-id").
		WithIdTokenFile("not-exist-file").
		SafeBuild()
	assert.NoError(t, err)

	req = request.NewHttpRequestBuilder().WithSigningAlgorithm(algorithm.GetDefaultSigningAlgorithm()).Build()
	_, err = credentials.ProcessAuthRequest(nil, req)
	assert.Error(t, err)
}

func TestGlobalCredentialsBuilder_SafeBuild(t *testing.T) {
	_, err := NewGlobalCredentialsBuilder().WithSk("").SafeBuild()
	assert.ErrorContains(t, err, "build credentials failed: input sk cannot be an empty string")

	credentials, err := NewGlobalCredentialsBuilder().
		WithAk("ak").
		WithSk("sk").
		WithSecurityToken("token").
		WithIdTokenFile("file").
		WithIdpId("idp-id").
		WithIamEndpointOverride("iam-endpoint").
		WithDomainId("domain-id").
		WithDerivedPredicate(func(httpRequest *request.DefaultHttpRequest) bool {
			return false
		}).
		SafeBuild()
	assert.NoError(t, err)
	assert.IsType(t, &GlobalCredentials{}, credentials)
	assert.Equal(t, "ak", credentials.AK)
	assert.Equal(t, "sk", credentials.SK)
	assert.Equal(t, "token", credentials.SecurityToken)
	assert.Equal(t, "idp-id", credentials.IdpId)
	assert.Equal(t, "domain-id", credentials.DomainId)
	assert.Equal(t, "iam-endpoint", credentials.IamEndpoint)
	assert.Equal(t, "file", credentials.IdTokenFile)
	assert.NotNil(t, credentials.DerivedPredicate)

	credentials.AK = "new_ak"
	credentials.DomainId = "new-domain-id"
	assert.Equal(t, "new_ak", credentials.AK)
	assert.Equal(t, "new-domain-id", credentials.DomainId)
}

// Deprecated
func TestGlobalCredentialsBuilder_Build(t *testing.T) {
	assert.NotPanics(t, func() {
		credentials := NewGlobalCredentialsBuilder().WithAk("ak").WithSk("sk").Build()
		assert.NotNil(t, credentials)
	})
}
