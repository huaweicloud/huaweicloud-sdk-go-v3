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
			builder.AddQueryParam(k, v)
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
				"Signature=eb70b3a0caf6633a8b8e4317622d8a0f6819821ec84896ceb7739a57e2d9ed76",
		},
		{
			TestParam: testParam2,
			Expected: "SDK-HMAC-SHA256 Access=AccessKey, SignedHeaders=x-sdk-date, " +
				"Signature=7f587e76874c5079f7e98769ab285ee75bea28e30110ab10dec4fc4d12c32f62",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			req := buildReqWithTestcase(c)
			result, err := signerInst.Sign(req, ak, sk)
			assert.Nil(t, err)
			assert.Equal(t, c.Expected, result["Authorization"])
		})
	}
}
