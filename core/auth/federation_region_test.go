// Copyright 2026 Shubham Padkonde
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package auth

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/internal"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/stretchr/testify/require"
)

type federationTransport func(*http.Request) (*http.Response, error)

func (f federationTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	return f(r)
}

func TestFederationRegionRequests(t *testing.T) {
	previous, present := os.LookupEnv(internal.IamEndpointEnv)
	require.NoError(t, os.Unsetenv(internal.IamEndpointEnv))
	t.Cleanup(func() {
		if present {
			require.NoError(t, os.Setenv(internal.IamEndpointEnv, previous))
		}
	})
	for _, kind := range []string{"basic", "global"} {
		for _, scope := range []string{"explicit", "cached", "discovered"} {
			t.Run(kind+"/"+scope, func(t *testing.T) {
				file := filepath.Join(t.TempDir(), "id-token")
				require.NoError(t, ioutil.WriteFile(file, []byte("id-token"), 0600))
				base := BaseCredentials{IdpId: t.Name(), IdTokenFile: file, StsAccessor: internal.NewFederalAccessor()}
				cacheKey := base.IdpId
				if kind == "basic" {
					cacheKey += "tr-west-1"
				}
				t.Cleanup(func() {
					cache := getCache()
					cache.mu.Lock()
					defer cache.mu.Unlock()
					delete(cache.data, cacheKey)
				})
				if scope == "cached" {
					getCache().put(cacheKey, "scope-id")
				}
				var credential ICredential
				var shared *BaseCredentials
				if kind == "basic" {
					c := &BasicCredentials{BaseCredentials: base}
					if scope == "explicit" {
						c.ProjectId = "scope-id"
					}
					credential, shared = c, &c.BaseCredentials
				} else {
					c := &GlobalCredentials{BaseCredentials: base}
					if scope == "explicit" {
						c.DomainId = "scope-id"
					}
					credential, shared = c, &c.BaseCredentials
				}
				var paths []string
				transport := federationTransport(func(r *http.Request) (*http.Response, error) {
					require.Equal(t, "https", r.URL.Scheme)
					require.Equal(t, "iam.tr-west-1.myhuaweicloud.com", r.URL.Host)
					paths = append(paths, r.URL.Path)
					header := http.Header{"Content-Type": []string{"application/json"}}
					header.Set("X-Request-Id", "request-id")
					header.Set("X-IAM-Trace-Id", "trace-id")
					var body string
					switch r.URL.Path {
					case "/v3.0/OS-AUTH/id-token/tokens":
						require.Equal(t, http.MethodPost, r.Method)
						header.Set("X-Subject-Token", "auth-token")
						body = `{"token":{"expires_at":"2099-01-01T00:00:00Z"}}`
					case "/v3.0/OS-CREDENTIAL/securitytokens":
						require.Equal(t, http.MethodPost, r.Method)
						body = `{"credential":{"access":"access","secret":"secret","securitytoken":"security-token","expires_at":"2099-01-01T00:00:00Z"}}`
					case "/v3/projects":
						require.Equal(t, "tr-west-1", r.URL.Query().Get("name"))
						body = `{"projects":[{"id":"scope-id"}]}`
					case "/v3/auth/domains":
						body = `{"domains":[{"id":"scope-id"}]}`
					default:
						t.Fatalf("unexpected request: %s", r.URL)
					}
					return &http.Response{StatusCode: 200, Header: header, Body: ioutil.NopCloser(strings.NewReader(body)), Request: r}, nil
				})
				client := impl.NewDefaultHttpClient(config.DefaultHttpConfig().WithHttpRoundTripper(transport))
				_, err := credential.ProcessAuthParams(client, "tr-west-1")
				require.NoError(t, err)
				credential.(IDerivedCredential).ProcessDerivedAuthParams("service", "tr-west-1")
				require.NoError(t, shared.ProcessSts(client))
				expected := []string{"/v3.0/OS-AUTH/id-token/tokens", "/v3.0/OS-CREDENTIAL/securitytokens"}
				if scope == "discovered" {
					if kind == "basic" {
						expected = append(expected, "/v3/projects")
					} else {
						expected = append(expected, "/v3/auth/domains")
					}
				}
				require.Equal(t, expected, paths)
				require.Equal(t, "security-token", shared.SecurityToken)
				shared.expireAt = 0
				require.NoError(t, shared.ProcessSts(client))
				expected = append(expected, "/v3.0/OS-AUTH/id-token/tokens", "/v3.0/OS-CREDENTIAL/securitytokens")
				require.Equal(t, expected, paths)
			})
		}
	}
}

type endpointRecordingAccessor struct {
	endpoints []string
}

func (a *endpointRecordingAccessor) GetCredential(options ...internal.StsAccessorOption) (*internal.Credential, error) {
	config := &internal.StsAccessorConfig{}
	for _, option := range options {
		option(config)
	}
	a.endpoints = append(a.endpoints, config.IamEndpoint)
	return &internal.Credential{
		Access: "access", Secret: "secret", SecurityToken: "token",
		ExpireAt: time.Now().Unix() + 10000,
	}, nil
}

func TestFederationRegionEndpoint(t *testing.T) {
	previous, present := os.LookupEnv(internal.IamEndpointEnv)
	t.Cleanup(func() {
		if present {
			require.NoError(t, os.Setenv(internal.IamEndpointEnv, previous))
		} else {
			require.NoError(t, os.Unsetenv(internal.IamEndpointEnv))
		}
	})
	cases := []struct {
		name, region, explicit, environment, expected string
	}{
		{"international", "tr-west-1", "", "", "https://iam.tr-west-1.myhuaweicloud.com"},
		{"china", "cn-north-4", "", "", "https://iam.cn-north-4.myhuaweicloud.com"},
		{"unknown", "unknown-region", "", "", internal.DefaultIamEndpoint},
		{"no region", "", "", "", internal.DefaultIamEndpoint},
		{"environment", "tr-west-1", "", "iam.example.test", "https://iam.example.test"},
		{"explicit", "tr-west-1", "https://explicit.example.test", "iam.example.test", "https://explicit.example.test"},
	}
	for _, kind := range []string{"basic", "global"} {
		for _, c := range cases {
			t.Run(kind+"/"+c.name, func(t *testing.T) {
				require.NoError(t, os.Setenv(internal.IamEndpointEnv, c.environment))
				accessor := &endpointRecordingAccessor{}
				base := BaseCredentials{IamEndpoint: c.explicit, StsAccessor: accessor}
				var credential ICredential
				var shared *BaseCredentials
				if kind == "basic" {
					basic := &BasicCredentials{BaseCredentials: base, ProjectId: "project"}
					credential, shared = basic, &basic.BaseCredentials
				} else {
					global := &GlobalCredentials{BaseCredentials: base, DomainId: "domain"}
					credential, shared = global, &global.BaseCredentials
				}
				_, err := credential.ProcessAuthParams(nil, c.region)
				require.NoError(t, err)
				credential.(IDerivedCredential).ProcessDerivedAuthParams("service", c.region)
				require.NoError(t, shared.ProcessSts(nil))
				require.Equal(t, []string{c.expected}, accessor.endpoints)
				// Cached credentials do not perform another exchange, but a refresh
				// must keep using the same regional endpoint.
				require.NoError(t, shared.ProcessSts(nil))
				require.Len(t, accessor.endpoints, 1)
				shared.expireAt = 0
				require.NoError(t, shared.ProcessSts(nil))
				require.Equal(t, []string{c.expected, c.expected}, accessor.endpoints)
			})
		}
	}
}
