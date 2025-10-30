// Copyright 2023 Huawei Technologies Co.,Ltd.
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

package signer

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

const (
	ak       = "AccessKey"
	sk       = "SecretKey"
	host     = "example.huaweicloud.com"
	endpoint = "https://" + host
	path     = "/path"
)

type TestBody struct {
	Name string
	Id   int
}

type TestParam struct {
	Name, Method, Endpoint, Path string
	Body                         interface{}
	Queries                      map[string]interface{}
	Headers                      map[string]string
}

type TestCase struct {
	TestParam
	Expected string
}

var (
	testParam1 = TestParam{
		Name:     "test1",
		Method:   "GET",
		Endpoint: endpoint,
		Path:     path,
		Body:     nil,
		Queries:  map[string]interface{}{"limit": 1},
		Headers:  map[string]string{"X-Sdk-Date": "20060102T150405Z", "TEST_UNDERSCORE": "TEST_VALUE"},
	}
	testParam2 = TestParam{
		Name:     "test2",
		Method:   "POST",
		Endpoint: endpoint,
		Path:     path,
		Body:     &TestBody{Name: "test", Id: 1},
		Queries:  map[string]interface{}{"key": "value"},
		Headers:  map[string]string{"X-Sdk-Date": "20060102T150405Z", "TEST_UNDERSCORE": "TEST_VALUE", "Content-Type": "application/json"},
	}
)

func buildReqWithTestcase(tc TestCase) *request.DefaultHttpRequest {
	builder := request.NewHttpRequestBuilder().WithMethod(tc.Method).WithEndpoint(tc.Endpoint).WithPath(tc.Path)
	if tc.Body != nil {
		builder.WithBody("body", tc.Body)
	}
	if tc.Queries != nil {
		for k, v := range tc.Queries {
			builder.AddQueryParam(k, reflect.ValueOf(v))
		}
	}
	if tc.Headers != nil {
		for k, v := range tc.Headers {
			builder.AddHeaderParam(k, v)
		}
	}
	return builder.Build()
}

func TestSigner_Sign(t *testing.T) {
	cases := []TestCase{
		{
			TestParam: testParam1,
			Expected: "SDK-HMAC-SHA256 Access=AccessKey, SignedHeaders=x-sdk-date, " +
				"Signature=5a2ce64c865e0e6046321c6f3d5a77ba8413eeaf355c3166c03d58d02ac79624",
		},
		{
			TestParam: testParam2,
			Expected: "SDK-HMAC-SHA256 Access=AccessKey, SignedHeaders=x-sdk-date, " +
				"Signature=cecc2af119b18ab70b4d094c0750f3b42c02f254903179a0fc2cc72fc9db4f59",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			req := buildReqWithTestcase(c)
			result, err := signerInst.Sign(req, ak, sk)
			assert.NoError(t, err)
			assert.Equal(t, c.Expected, result["Authorization"])
		})
	}
}

func TestSigner_canonicalQueryString(t *testing.T) {
	req := request.NewHttpRequestBuilder().
		AddQueryParam("limit", reflect.ValueOf(1)).
		AddQueryParam("enable", reflect.ValueOf(true)).
		AddQueryParam("test", reflect.ValueOf("一 (&=?!#%.*)")).
		AddQueryParam("path", reflect.ValueOf("/tmp/123")).
		AddQueryParam("multi", reflect.ValueOf([]int{1, 2, 3})).
		Build()
	str := canonicalQueryString(req)
	assert.Equal(t, "enable=true&limit=1&multi=1&multi=2&multi=3&path=%2Ftmp%2F123&test=%E4%B8%80%20%28%26%3D%3F%21%23%25.%2A%29", str)
}
