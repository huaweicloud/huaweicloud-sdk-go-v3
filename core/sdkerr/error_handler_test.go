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

package sdkerr

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/response"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestDefaultErrorHandler_HandleError(t *testing.T) {
	resp := &http.Response{StatusCode: 200}
	httpResp := &response.DefaultHttpResponse{Response: resp}
	err := DefaultErrorHandler{}.HandleError(nil, httpResp)
	assert.NoError(t, err)

	resp = &http.Response{StatusCode: 400, Body: io.NopCloser(strings.NewReader("some errors"))}
	httpResp = &response.DefaultHttpResponse{Response: resp}
	err = DefaultErrorHandler{}.HandleError(nil, httpResp)
	assert.ErrorContains(t, err, "some errors")
}
