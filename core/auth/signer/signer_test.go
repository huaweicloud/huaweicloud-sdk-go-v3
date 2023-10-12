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
	ak   = "AccessKey"
	sk   = "SecretKey"
	host = "example.huaweicloud.com"
)

type testBody struct {
	Name string
	Id   int
}

func TestSigner_Sign(t *testing.T) {
	req := request.NewHttpRequestBuilder().
		WithMethod("GET").
		WithEndpoint("https://"+host).
		WithPath("/path").
		AddHeaderParam("X-Sdk-Date", "20060102T150405Z").
		AddQueryParam("limit", 1).
		Build()
	result, err := signerInst.Sign(req, ak, sk)
	assert.Nil(t, err)
	assert.Equal(t, "SDK-HMAC-SHA256 Access=AccessKey, SignedHeaders=x-sdk-date, "+
		"Signature=eb70b3a0caf6633a8b8e4317622d8a0f6819821ec84896ceb7739a57e2d9ed76", result["Authorization"])
}

func TestSigner_Sign2(t *testing.T) {
	body := &testBody{
		Name: "test",
		Id:   1,
	}
	req := request.NewHttpRequestBuilder().
		WithMethod("POST").
		WithEndpoint("https://"+host).
		WithPath("/path").
		WithBody("body", body).AddQueryParam("key", "value").
		AddHeaderParam("Content-Type", "application/json").
		AddHeaderParam("X-Sdk-Date", "20060102T150405Z").
		Build()
	result, err := signerInst.Sign(req, ak, sk)
	assert.Nil(t, err)
	assert.Equal(t, "SDK-HMAC-SHA256 Access=AccessKey, SignedHeaders=x-sdk-date, "+
		"Signature=7f587e76874c5079f7e98769ab285ee75bea28e30110ab10dec4fc4d12c32f62", result["Authorization"])
}
