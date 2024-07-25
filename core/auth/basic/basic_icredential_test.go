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
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCredentials_NeedUpdate(t *testing.T) {
	credentials, err := NewCredentialsBuilder().SafeBuild()
	assert.Nil(t, err)
	// Manually specifying a security token
	credentials, err = NewCredentialsBuilder().
		WithAk("ak").
		WithSk("sk").
		WithSecurityToken("token").
		SafeBuild()
	assert.Nil(t, err)
	assert.False(t, credentials.needUpdateSecurityToken())
	// Automatically update the expired security token
	credentials.expiredAt = 1
	assert.True(t, credentials.needUpdateSecurityToken())
	// The security token has not expired
	credentials.expiredAt = time.Now().Unix() + 120
	assert.False(t, credentials.needUpdateSecurityToken())
}

func TestCredentials_NeedUpdateAuthToken(t *testing.T) {
	credentials, err := NewCredentialsBuilder().SafeBuild()
	assert.Nil(t, err)
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

func TestCredentialsBuilder_SafeBuild(t *testing.T) {
	cases := []struct {
		Name, IdpId, IdTokenFile, ProjectId, ExpectedErr string
	}{
		{"ProjectId missed", "id", "file", "", "ProjectId is required when using IdpId&IdTokenFile"},
		{"IdTokenFile missed", "id", "", "projectId", "IdTokenFile is required when using IdpId&IdTokenFile"},
		{"IdpId missed", "", "file", "projectId", "IdpId is required when using IdpId&IdTokenFile"},
		{"valid params", "id", "file", "projectId", ""},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			cred, err := NewCredentialsBuilder().
				WithIdpId(c.IdpId).
				WithIdTokenFile(c.IdTokenFile).
				WithProjectId(c.ProjectId).
				SafeBuild()
			var cte *sdkerr.CredentialsTypeError
			if errors.As(err, &cte) {
				assert.Equal(t, c.ExpectedErr, cte.ErrorMessage)
			} else {
				assert.Equal(t, c.IdpId, cred.IdpId)
				assert.Equal(t, c.IdTokenFile, cred.IdTokenFile)
				assert.Equal(t, c.ProjectId, cred.ProjectId)
			}
		})
	}
}

func TestCredentialsBuilder_SafeBuild2(t *testing.T) {
	cases := []struct {
		Name, Ak, Sk     string
		ExpectedContains []string
	}{
		{"Empty AK", "", "sk", []string{"input ak cannot be an empty string"}},
		{"Empty SK", "ak", "", []string{"input sk cannot be an empty string"}},
		{"Empty AK&SK", "", "", []string{"input ak cannot be an empty string", "input sk cannot be an empty string"}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			_, err := NewCredentialsBuilder().WithAk(c.Ak).WithSk(c.Sk).SafeBuild()
			assert.IsType(t, err, &sdkerr.CredentialsTypeError{})
			msg := err.(*sdkerr.CredentialsTypeError).ErrorMessage
			for _, contain := range c.ExpectedContains {
				assert.Contains(t, msg, contain)
			}
		})
	}

	cases2 := []struct {
		Name string
		Func func() (*Credentials, error)
	}{
		{"Valid AK&SK", func() (*Credentials, error) {
			return NewCredentialsBuilder().
				WithAk("ak").
				WithSk("sk").
				SafeBuild()
		}},
		{"Invalid AK&SK then correct it", func() (*Credentials, error) {
			return NewCredentialsBuilder().
				WithAk("").
				WithSk("").
				WithAk("ak").
				WithSk("sk").
				SafeBuild()
		}},
	}

	for _, c := range cases2 {
		t.Run(c.Name, func(t *testing.T) {
			credentials, err := c.Func()
			assert.Nil(t, err)
			assert.NotEmpty(t, credentials.AK)
			assert.NotEmpty(t, credentials.SK)
		})
	}
}

func TestCredentials_ProcessAuthParams(t *testing.T) {
	cases := []struct {
		Name, RegionId, Data, ExpectedProjectId string
		ExpectedError                           error
	}{
		{"One Project", "region-id-1", "{\"projects\":[{\"id\":\"project_id\"}]}", "project_id", nil},
		{"No Project", "region-id-2", "{\"projects\":[]}", "",
			errors.New("failed to get project id of region 'region-id-2' automatically," +
				" X-IAM-Trace-Id=trace-id, confirm that the project exists in your account," +
				" or set project id manually:" +
				" basic.NewCredentialsBuilder().WithAk(ak).WithSk(sk).WithProjectId(projectId).SafeBuild()")},
		{"Multi Projects", "region-id-3", "{\"projects\":[{\"id\":\"project_id1\"},{\"id\":\"project_id2\"}]}", "",
			errors.New("multiple project ids found: [project_id1,project_id2]," +
				" X-IAM-Trace-Id=trace-id, please select one when initializing the credentials:" +
				" basic.NewCredentialsBuilder().WithAk(ak).WithSk(sk).WithProjectId(projectId).SafeBuild()")},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.Header().Set("X-IAM-Trace-Id", "trace-id")
				_, err := fmt.Fprintln(w, c.Data)
				assert.Nil(t, err)
			}))
			defer ts.Close()

			credentials, err := NewCredentialsBuilder().
				WithAk("ak").
				WithSk("sk").
				WithIamEndpointOverride(ts.URL).
				SafeBuild()
			assert.Nil(t, err)

			defer func() {
				if c.ExpectedError != nil {
					assert.Equal(t, c.ExpectedError, recover())
				} else {
					assert.Equal(t, c.ExpectedProjectId, credentials.ProjectId)
				}
			}()
			client := impl.NewDefaultHttpClient(config.DefaultHttpConfig())
			credentials.ProcessAuthParams(client, c.RegionId)
		})
	}
}
