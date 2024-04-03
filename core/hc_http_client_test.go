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

package core

import (
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/exchange"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testStruct struct {
	HttpStatusCode int `json:"-"`
}

func TestHcHttpClient_SyncInvokeWithExtraHeaders(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "{}")
		assert.Nil(t, err)
		// assert custom user-agent
		assert.Equal(t, "huaweicloud-usdk-go/3.0;test-user-agent", r.Header.Get("User-Agent"))
		// assert client header
		assert.Equal(t, "test-value-1", r.Header.Get("test-header-1"))
		// assert request header override client header
		assert.Equal(t, "test-value-override", r.Header.Get("test-header-2"))
		// assert request header
		assert.Equal(t, "test-value-3", r.Header.Get("test-header-3"))
	}))
	defer ts.Close()

	client, err := NewHcHttpClientBuilder().
		WithHttpConfig(config.DefaultHttpConfig()).
		WithEndpoints([]string{ts.URL}).
		WithCredential(&basic.Credentials{AK: "test", SK: "test"}).
		SafeBuild()
	assert.Nil(t, err)
	client.extraHeaders = map[string]string{
		"User-Agent":    "test-user-agent",
		"test-header-1": "test-value-1",
		"test-header-2": "test-value-2"}

	reqDef := def.NewHttpRequestDefBuilder().
		WithMethod("GET").WithPath("/").
		WithContentType("application/json").
		WithResponse(&testStruct{}).
		Build()
	req := &testStruct{}
	exc := &exchange.SdkExchange{ApiReference: &exchange.ApiReference{}}
	extraHeaders := map[string]string{
		"test-header-2": "test-value-override",
		"test-header-3": "test-value-3"}
	resp, err := client.SyncInvokeWithExtraHeaders(req, reqDef, exc, extraHeaders)
	assert.Nil(t, err)
	s, ok := resp.(*testStruct)
	assert.True(t, ok)
	assert.Equal(t, 200, s.HttpStatusCode)
}
