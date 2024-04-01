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

package sdkerr

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

func mockHeader() http.Header {
	header := http.Header{}
	header.Add("X-Request-Id", "97e2***11df")
	return header
}

func TestNewServiceResponseError(t *testing.T) {
	data := []byte("{\"error_code\":\"XXX.0001\"," +
		"\"error_msg\":\"Some errors occurred.\"," +
		"\"encoded_authorization_message\":\"Egpjbi1ub*****GoxMzgrcA==\"}")
	response := &http.Response{
		StatusCode: 400,
		Header:     mockHeader(),
		Body:       ioutil.NopCloser(bytes.NewBuffer(data)),
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.Nil(t, err)
	}(response.Body)

	responseError := NewServiceResponseError(response)
	assert.Equal(t, 400, responseError.StatusCode)
	assert.Equal(t, "XXX.0001", responseError.ErrorCode)
	assert.Equal(t, "Some errors occurred.", responseError.ErrorMessage)
	assert.Equal(t, "97e2***11df", responseError.RequestId)
}

func TestNewServiceResponseError2(t *testing.T) {
	data := []byte("{\"error\":" +
		"{\"code\":\"XXX.0001\"," +
		"\"message\":\"Some errors occurred.\"," +
		"\"encoded_authorization_message\":\"Egpjbi1ub*****GoxMzgrcA==\"}}")
	response := &http.Response{
		StatusCode: 400,
		Header:     mockHeader(),
		Body:       ioutil.NopCloser(bytes.NewBuffer(data)),
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.Nil(t, err)
	}(response.Body)

	responseError := NewServiceResponseError(response)
	assert.Equal(t, 400, responseError.StatusCode)
	assert.Equal(t, "XXX.0001", responseError.ErrorCode)
	assert.Equal(t, "Some errors occurred.", responseError.ErrorMessage)
	assert.Equal(t, "97e2***11df", responseError.RequestId)
}

func TestNewServiceResponseError3(t *testing.T) {
	data := []byte("{\"error\":" +
		"{\"code\":\"XXX.0001\"," +
		"\"message\":\"Some errors occurred.\"}," +
		"\"error_code\":\"XXX.0001\"," +
		"\"error_msg\":\"Some errors occurred.\"," +
		"\"encoded_authorization_message\":\"Egpjbi1ub*****GoxMzgrcA==\"}")
	response := &http.Response{
		StatusCode: 400,
		Header:     mockHeader(),
		Body:       ioutil.NopCloser(bytes.NewBuffer(data)),
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.Nil(t, err)
	}(response.Body)

	responseError := NewServiceResponseError(response)
	assert.Equal(t, 400, responseError.StatusCode)
	assert.Equal(t, "XXX.0001", responseError.ErrorCode)
	assert.Equal(t, "Some errors occurred.", responseError.ErrorMessage)
	assert.Equal(t, "97e2***11df", responseError.RequestId)
}

func TestNewServiceResponseError4(t *testing.T) {
	data := []byte("{\"invalid_key\":\"invalid_value\"}")
	response := &http.Response{
		StatusCode: 400,
		Header:     mockHeader(),
		Body:       ioutil.NopCloser(bytes.NewBuffer(data)),
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.Nil(t, err)
	}(response.Body)

	responseError := NewServiceResponseError(response)
	assert.Equal(t, 400, responseError.StatusCode)
	assert.Equal(t, "", responseError.ErrorCode)
	assert.Equal(t, "{\"invalid_key\":\"invalid_value\"}", responseError.ErrorMessage)
	assert.Equal(t, "97e2***11df", responseError.RequestId)
}

func TestNewServiceResponseError5(t *testing.T) {
	data := []byte("invalid json data")
	response := &http.Response{
		StatusCode: 400,
		Header:     mockHeader(),
		Body:       ioutil.NopCloser(bytes.NewBuffer(data)),
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		assert.Nil(t, err)
	}(response.Body)

	responseError := NewServiceResponseError(response)
	assert.Equal(t, 400, responseError.StatusCode)
	assert.Equal(t, "", responseError.ErrorCode)
	assert.Equal(t, "invalid json data", responseError.ErrorMessage)
	assert.Equal(t, "97e2***11df", responseError.RequestId)
}
