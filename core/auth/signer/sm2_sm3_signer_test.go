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
	"strings"
	"testing"
)

func TestSM2SM3Signer_Sign(t *testing.T) {
	req := request.NewHttpRequestBuilder().
		WithMethod("GET").
		WithEndpoint("https://"+host).
		WithPath("/path").
		AddHeaderParam("X-Sdk-Date", "20060102T150405Z").
		AddQueryParam("limit", 1).
		Build()
	result, err := sm2sm3SignerInst.Sign(req, ak, sk)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(result["Authorization"], "SDK-SM2-SM3"))
}

func TestSM2SM3Signer_Sign2(t *testing.T) {
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
	result, err := sm2sm3SignerInst.Sign(req, ak, sk)
	assert.Nil(t, err)
	assert.True(t, strings.HasPrefix(result["Authorization"], "SDK-SM2-SM3"))
}
