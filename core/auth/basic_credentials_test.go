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

func TestBasicCredentials_ProcessAuthParams(t *testing.T) {
	cases := []struct {
		Name, RegionId, Data, ExpectedProjectId string
		StatusCode                              int
		ExpectedError                           string
	}{
		{"Bad Request", "region-id-1", "{\"error_code\":\"XXX.001\",\"error_msg\":\"Bad Request\"}", "", 400,
			"failed to get project id of region 'region-id-1' automatically, {\"status_code\":400,\"request_id\":\"\",\"error_code\":\"XXX.001\",\"error_message\":\"Bad Request, X-IAM-Trace-Id=trace-id\",\"encoded_authorization_message\":\"\"}"},
		{"One Project", "region-id-1", "{\"projects\":[{\"id\":\"project_id\"}]}", "project_id", 200, ""},
		{"No Project", "region-id-2", "{\"projects\":[]}", "", 200,
			"failed to get project id of region 'region-id-2' automatically," +
				" X-IAM-Trace-Id=trace-id, confirm that the project exists in your account," +
				" or set project id manually:" +
				" basic.NewCredentialsBuilder().WithAk(ak).WithSk(sk).WithProjectId(projectId).SafeBuild()"},
		{"Multi Projects", "region-id-3", "{\"projects\":[{\"id\":\"project_id1\"},{\"id\":\"project_id2\"}]}", "", 200,
			"multiple project ids found: [project_id1,project_id2]," +
				" X-IAM-Trace-Id=trace-id, please select one when initializing the credentials:" +
				" basic.NewCredentialsBuilder().WithAk(ak).WithSk(sk).WithProjectId(projectId).SafeBuild()"},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("X-IAM-Trace-Id", "trace-id")
				w.WriteHeader(c.StatusCode)
				_, err := fmt.Fprintln(w, c.Data)
				assert.NoError(t, err)
			}))
			defer ts.Close()

			credentials, err := NewBasicCredentialsBuilder().
				WithAk("ak").
				WithSk("sk").
				WithIamEndpointOverride(ts.URL).
				SafeBuild()
			assert.NoError(t, err)

			defer func() {
				if c.ExpectedError != "" {
					assert.Equal(t, c.ExpectedError, strings.TrimSpace(fmt.Sprintf("%v", recover())))
				} else {
					assert.Equal(t, c.ExpectedProjectId, credentials.ProjectId)
				}
			}()
			client := impl.NewDefaultHttpClient(config.DefaultHttpConfig())
			credentials.ProcessAuthParams(client, c.RegionId)
		})
	}
}

func TestBasicCredentials_ProcessAuthParams2(t *testing.T) {
	credentials, err := NewBasicCredentialsBuilder().WithAk("ak").WithSk("sk").SafeBuild()
	assert.NoError(t, err)
	client := impl.NewDefaultHttpClient(config.DefaultHttpConfig().WithSigningAlgorithm("test"))
	assert.Panics(t, func() {
		credentials.ProcessAuthParams(client, "region")
	})

	credentials, err = NewBasicCredentialsBuilder().WithAk("ak").WithSk("sk").SafeBuild()
	assert.NoError(t, err)
	credentials.ProjectId = "project-id"
	credentials, ok := credentials.ProcessAuthParams(nil, "").(*BasicCredentials)
	assert.True(t, ok)
	assert.Equal(t, "project-id", credentials.ProjectId)

	credentials, err = NewBasicCredentialsBuilder().WithAk("ak").WithSk("sk").SafeBuild()
	assert.NoError(t, err)
	getCache().put("akregion", "id")
	credentials, ok = credentials.ProcessAuthParams(nil, "region").(*BasicCredentials)
	assert.True(t, ok)
	assert.Equal(t, "id", credentials.ProjectId)
}

func TestBasicCredentials_FederalAuth(t *testing.T) {
	var count int32
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	filePath := filepath.Join(tmpDir, "BasicTestIdTokenFile")
	err := os.WriteFile(filePath, []byte("id_token"), 0600)
	filePath, err = filepath.Abs(filePath)
	assert.NoError(t, err)

	cred, err := NewBasicCredentialsBuilder().
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

func TestBasicCredentials_ProcessDerivedAuthParams(t *testing.T) {
	credentials, err := NewBasicCredentialsBuilder().WithAk("ak").WithSk("sk").SafeBuild()
	assert.NoError(t, err)
	credentials.ProcessDerivedAuthParams("name", "region")
	assert.Equal(t, "name", credentials.derivedAuthServiceName)
	assert.Equal(t, "region", credentials.regionId)
}

func TestBasicCredentials_ProcessAuthRequest(t *testing.T) {
	credentials, err := NewBasicCredentialsBuilder().
		WithAk("ak").
		WithSk("sk").
		WithSecurityToken("xxx").
		WithProjectId("project-id").
		SafeBuild()
	assert.NoError(t, err)
	req := request.NewHttpRequestBuilder().WithSigningAlgorithm(algorithm.GetDefaultSigningAlgorithm()).Build()
	req, err = credentials.ProcessAuthRequest(nil, req)
	assert.NoError(t, err)
	assert.Equal(t, req.GetHeaderParams()["X-Project-Id"], "project-id")
	assert.Contains(t, req.GetHeaderParams(), "Authorization")

	credentials, err = NewBasicCredentialsBuilder().
		WithIdpId("idp-id").
		WithIdTokenFile("not-exist-file").
		SafeBuild()
	assert.NoError(t, err)

	req = request.NewHttpRequestBuilder().WithSigningAlgorithm(algorithm.GetDefaultSigningAlgorithm()).Build()
	_, err = credentials.ProcessAuthRequest(nil, req)
	assert.Error(t, err)
}

func TestBasicCredentialsBuilder_SafeBuild(t *testing.T) {
	_, err := NewBasicCredentialsBuilder().WithAk("").SafeBuild()
	assert.ErrorContains(t, err, "build credentials failed: input ak cannot be an empty string")

	credentials, err := NewBasicCredentialsBuilder().
		WithAk("ak").
		WithSk("sk").
		WithSecurityToken("token").
		WithIdTokenFile("file").
		WithIdpId("idp-id").
		WithIamEndpointOverride("iam-endpoint").
		WithProjectId("project-id").
		WithDerivedPredicate(func(httpRequest *request.DefaultHttpRequest) bool {
			return false
		}).
		SafeBuild()
	assert.NoError(t, err)
	assert.IsType(t, &BasicCredentials{}, credentials)
	assert.Equal(t, "ak", credentials.AK)
	assert.Equal(t, "sk", credentials.SK)
	assert.Equal(t, "token", credentials.SecurityToken)
	assert.Equal(t, "idp-id", credentials.IdpId)
	assert.Equal(t, "project-id", credentials.ProjectId)
	assert.Equal(t, "iam-endpoint", credentials.IamEndpoint)
	assert.Equal(t, "file", credentials.IdTokenFile)
	assert.NotNil(t, credentials.DerivedPredicate)

	credentials.AK = "new_ak"
	credentials.ProjectId = "new-project-id"
	assert.Equal(t, "new_ak", credentials.AK)
	assert.Equal(t, "new-project-id", credentials.ProjectId)
}

// Deprecated
func TestBasicCredentialsBuilder_Build(t *testing.T) {
	assert.NotPanics(t, func() {
		credentials := NewBasicCredentialsBuilder().WithAk("ak").WithSk("sk").Build()
		assert.NotNil(t, credentials)
	})
}
