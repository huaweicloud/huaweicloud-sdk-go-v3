// Copyright 2020 Huawei Technologies Co.,Ltd.
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
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"os"
	"reflect"
)

type HcHttpClientBuilder struct {
	credentialsType string
	credentials     auth.ICredential
	endpoint        string
	httpConfig      *config.HttpConfig
}

func NewHcHttpClientBuilder() *HcHttpClientBuilder {
	hcHttpClientBuilder := &HcHttpClientBuilder{
		credentialsType: "basic.Credentials",
	}
	return hcHttpClientBuilder
}

func (builder *HcHttpClientBuilder) WithCredentialsType(credentialsType string) *HcHttpClientBuilder {
	builder.credentialsType = credentialsType
	return builder
}

func (builder *HcHttpClientBuilder) WithEndpoint(endpoint string) *HcHttpClientBuilder {
	builder.endpoint = endpoint
	return builder
}

func (builder *HcHttpClientBuilder) WithHttpConfig(httpConfig *config.HttpConfig) *HcHttpClientBuilder {
	builder.httpConfig = httpConfig
	return builder
}

func (builder *HcHttpClientBuilder) WithCredential(iCredential auth.ICredential) *HcHttpClientBuilder {
	builder.credentials = iCredential
	return builder
}

func (builder *HcHttpClientBuilder) Build() *HcHttpClient {
	if builder.httpConfig == nil {
		builder.httpConfig = config.DefaultHttpConfig()
	}

	if builder.credentials == nil {
		builder.LoadCredentialFromEnv()
	}

	if reflect.TypeOf(builder.credentials).String() != builder.credentialsType {
		panic(fmt.Sprintf("Need credential type is %s, actually is %s", builder.credentialsType, reflect.TypeOf(builder.credentials).String()))
	}

	defaultHttpClient := impl.NewDefaultHttpClient(builder.httpConfig)

	hcHttpClient := NewHcHttpClient(defaultHttpClient).WithEndpoint(builder.endpoint).WithCredential(builder.credentials)
	return hcHttpClient
}

func (builder *HcHttpClientBuilder) LoadCredentialFromEnv() {
	ak := os.Getenv("HUAWEICLOUD_SDK_AK")
	sk := os.Getenv("HUAWEICLOUD_SDK_SK")
	projectId := os.Getenv("HUAWEICLOUD_SDK_PROJECT_ID")
	domainId := os.Getenv("HUAWEICLOUD_SDK_DOMAIN_ID")
	if builder.credentialsType == "basic.Credentials" {
		builder.credentials = basic.NewCredentialsBuilder().
			WithAk(ak).
			WithSk(sk).
			WithProjectId(projectId).
			Build()
	}
	if builder.credentialsType == "global.Credentials" {
		builder.credentials = global.NewCredentialsBuilder().
			WithAk(ak).
			WithSk(sk).
			WithDomainId(domainId).
			Build()
	}
}
