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
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/exchange"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type QueryRequest struct {
	Marker  *string    `json:"marker,omitempty"`
	Limit   *int32     `json:"limit,omitempty"`
	Id      *[]int32   `json:"id,omitempty"`
	SortKey *[]string  `json:"sort_key,omitempty"`
	SortDir *[]SortDir `json:"sort_dir,omitempty"`
}

type SortDir struct {
	value string
}

type SortDirEnum struct {
	ASC  SortDir
	DESC SortDir
}

func GetSortDirEnum() SortDirEnum {
	return SortDirEnum{
		ASC: SortDir{
			value: "asc",
		},
		DESC: SortDir{
			value: "desc",
		},
	}
}

func (c SortDir) Value() string {
	return c.value
}

func (c SortDir) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *SortDir) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

func GenReqDefForQueryRequest() *def.HttpRequestDef {
	reqDefBuilder := def.NewHttpRequestDefBuilder().
		WithMethod(http.MethodGet).
		WithPath("/").
		WithResponse(new(testStruct)).
		WithContentType("application/json")

	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Limit").
		WithJsonTag("limit").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Marker").
		WithJsonTag("marker").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("Id").
		WithJsonTag("id").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SortKey").
		WithJsonTag("sort_key").
		WithLocationType(def.Query))
	reqDefBuilder.WithRequestField(def.NewFieldDef().
		WithName("SortDir").
		WithJsonTag("sort_dir").
		WithLocationType(def.Query))

	requestDef := reqDefBuilder.Build()
	return requestDef
}

func TestHcHttpClient_SyncInvokeWithQuery(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "{}")
		assert.Nil(t, err)
		fmt.Println(r.URL.RawQuery)
	}))
	defer ts.Close()

	credentials, err := basic.NewCredentialsBuilder().
		WithAk("test").
		WithSk("test").
		WithProjectId("project-id").
		SafeBuild()
	assert.NoError(t, err)
	client, err := NewHcHttpClientBuilder().
		WithEndpoints([]string{ts.URL}).
		WithCredential(credentials).
		WithHttpConfig(config.DefaultHttpConfig().WithIgnoreSSLVerification(true)).
		SafeBuild()
	assert.Nil(t, err)

	reqDef := GenReqDefForQueryRequest()

	marker := "mark"
	limit := int32(10)
	id := []int32{100, 200}
	sortKey := []string{"a", "b"}
	sortDir := []SortDir{GetSortDirEnum().ASC, GetSortDirEnum().DESC}
	req := &QueryRequest{
		Marker:  &marker,
		Limit:   &limit,
		Id:      &id,
		SortKey: &sortKey,
		SortDir: &sortDir,
	}
	exc := &exchange.SdkExchange{
		ApiReference: &exchange.ApiReference{},
		Attributes:   make(map[string]interface{}),
	}
	resp, err := client.SyncInvoke(req, reqDef, exc)
	assert.Nil(t, err)
	s, ok := resp.(*testStruct)
	assert.True(t, ok)
	assert.Equal(t, 200, s.HttpStatusCode)
}

type testStruct struct {
	HttpStatusCode int `json:"-"`
}

func TestHcHttpClient_SyncInvokeWithExtraHeaders(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "{}")
		assert.NoError(t, err)
		// assert custom user-agent
		ua := r.Header.Get("User-Agent")
		assert.True(t, strings.Count(ua, ";") >= 2)
		assert.True(t, strings.HasPrefix(ua, "huaweicloud-usdk-go/3.0"))
		assert.True(t, strings.HasSuffix(ua, "test-user-agent"))
		// assert client header
		assert.Equal(t, "test-value-1", r.Header.Get("test-header-1"))
		// assert request header override client header
		assert.Equal(t, "test-value-override", r.Header.Get("test-header-2"))
		// assert request header
		assert.Equal(t, "test-value-3", r.Header.Get("test-header-3"))
	}))
	defer ts.Close()

	credentials, err := basic.NewCredentialsBuilder().WithAk("test").WithSk("test").SafeBuild()
	assert.NoError(t, err)

	client, err := NewHcHttpClientBuilder().
		WithEndpoints([]string{ts.URL}).
		WithCredential(credentials).
		WithHttpConfig(config.DefaultHttpConfig().WithIgnoreSSLVerification(true)).
		SafeBuild()
	assert.NoError(t, err)
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
	assert.NoError(t, err)
	s, ok := resp.(*testStruct)
	assert.True(t, ok)
	assert.Equal(t, 200, s.HttpStatusCode)
}

func TestHcHttpClient_Sync(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "{}")
		assert.Nil(t, err)
		// assert custom user-agent
		ua := r.Header.Get("User-Agent")
		assert.Equal(t, "huaweicloud-usdk-go/3.0; custom-user-agent", ua)
	}))
	defer ts.Close()

	credentials, err := basic.NewCredentialsBuilder().WithAk("test").WithSk("test").SafeBuild()
	assert.NoError(t, err)
	client, err := NewHcHttpClientBuilder().
		WithHttpConfig(config.DefaultHttpConfig().WithUserAgent("custom-user-agent").WithIgnoreSSLVerification(true)).
		WithEndpoints([]string{ts.URL}).
		WithCredential(credentials).
		SafeBuild()
	assert.Nil(t, err)

	reqDef := def.NewHttpRequestDefBuilder().
		WithMethod("GET").WithPath("/").
		WithContentType("application/json").
		WithResponse(&testStruct{}).
		Build()
	req := &testStruct{}
	resp, err := client.Sync(req, reqDef)
	assert.Nil(t, err)
	s, ok := resp.(*testStruct)
	assert.True(t, ok)
	assert.Equal(t, 200, s.HttpStatusCode)
}
