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

package internal

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/stretchr/testify/assert"
	"net"
	"net/http"
	"os"
	"testing"
)

type dnsErrorRoundTripper struct{}

func (t *dnsErrorRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, &net.DNSError{
		Name:        req.URL.Hostname(),
		Server:      "endpoint",
		Err:         "no such host",
		IsTimeout:   false,
		IsTemporary: false,
		IsNotFound:  true,
	}
}

func TestGetStsEndpointById(t *testing.T) {
	assert.Empty(t, GetStsEndpointById("region"))
	err := os.Setenv("HUAWEICLOUD_SDK_STS_ENDPOINT", "endpoint")
	assert.NoError(t, err)
	assert.Equal(t, "https://endpoint", GetStsEndpointById(""))
}

func TestGetAccountIdFromCallerIdentity(t *testing.T) {
	httpConfig := config.DefaultHttpConfig().WithHttpRoundTripper(&dnsErrorRoundTripper{})
	req := GetCallerIdentityRequest("endpoint", *httpConfig)
	client := impl.NewDefaultHttpClient(httpConfig)
	_, err := GetAccountIdFromCallerIdentity(client, req)
	assert.ErrorContains(t, err, "failed to get domain id from endpoint")
	assert.ErrorContains(t, err, "no such host")
}
