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
	service  = "demo"
	regionId = "test-region-1"
)

func TestDerivedSigner_Sign(t *testing.T) {
	req := request.NewHttpRequestBuilder().
		WithMethod("GET").
		WithEndpoint("https://"+host).
		WithPath("/path").
		AddHeaderParam("X-Sdk-Date", "20060102T150405Z").
		AddHeaderParam("TEST_UNDERSCORE", "TEST_VALUE").
		AddQueryParam("limit", 1).
		Build()
	result, err := GetDerivedSigner().Sign(req, ak, sk, service, regionId)
	assert.Nil(t, err)
	assert.Equal(t, "V11-HMAC-SHA256 Credential=AccessKey/20060102/test-region-1/demo,"+
		" SignedHeaders=x-sdk-date,"+
		" Signature=5579214897b4b5def742b24620946119573bccdfe7b2ea7c720e64f1a0944d07", result["Authorization"])
}

func TestDerivedSigner_Sign2(t *testing.T) {
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
		AddHeaderParam("TEST_UNDERSCORE", "TEST_VALUE").
		Build()
	result, err := GetDerivedSigner().Sign(req, ak, sk, service, regionId)
	assert.Nil(t, err)
	assert.Equal(t, "V11-HMAC-SHA256 Credential=AccessKey/20060102/test-region-1/demo, "+
		"SignedHeaders=x-sdk-date, "+
		"Signature=638f6f0deca2285bf3af0dfd7b186b447e60ed8611d4c6c93537aa8711b9bce9", result["Authorization"])
}

func TestDerivedSigner_Sign3(t *testing.T) {
	req := request.NewHttpRequestBuilder().
		WithMethod("GET").
		WithEndpoint("https://"+host).
		WithPath("/path").
		AddHeaderParam("X-Sdk-Date", "20060102T150405Z").
		AddHeaderParam("TEST_UNDERSCORE", "TEST_VALUE").
		AddQueryParam("limit", 1).
		Build()
	result, err := SignDerived(req, ak, sk, service, regionId)
	assert.Nil(t, err)
	assert.Equal(t, "V11-HMAC-SHA256 Credential=AccessKey/20060102/test-region-1/demo,"+
		" SignedHeaders=x-sdk-date,"+
		" Signature=5579214897b4b5def742b24620946119573bccdfe7b2ea7c720e64f1a0944d07", result["Authorization"])
}
