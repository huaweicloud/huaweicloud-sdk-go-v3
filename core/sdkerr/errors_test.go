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
	cases := []struct {
		Name                                                            string
		Data                                                            []byte
		StatusCode                                                      int
		ErrorCode, ErrorMessage, RequestId, EncodedAuthorizationMessage string
	}{
		{
			"error_code&error_msg",
			[]byte("{\"error_code\":\"XXX.0001\",\"error_msg\":\"Some errors occurred.\"," +
				"\"encoded_authorization_message\":\"Egpjbi1ub*****GoxMzgrcA==\"}"),
			400, "XXX.0001", "Some errors occurred.", "97e2***11df", "Egpjbi1ub*****GoxMzgrcA==",
		},
		{
			"code&message",
			[]byte("{\"error\":{\"code\":\"XXX.0001\",\"message\":\"Some errors occurred.\"," +
				"\"encoded_authorization_message\":\"Egpjbi1ub*****GoxMzgrcA==\"}}"),
			400, "XXX.0001", "Some errors occurred.", "97e2***11df", "Egpjbi1ub*****GoxMzgrcA==",
		},
		{
			"error_code&error_msg&code&message",
			[]byte("{\"error\":" +
				"{\"code\":\"XXX.0002\"," +
				"\"message\":\"Some errors occurred...\"}," +
				"\"error_code\":\"XXX.0001\"," +
				"\"error_msg\":\"Some errors occurred.\"," +
				"\"encoded_authorization_message\":\"Egpjbi1ub*****GoxMzgrcA==\"}"),
			400, "XXX.0001", "Some errors occurred.", "97e2***11df", "Egpjbi1ub*****GoxMzgrcA==",
		},
		{
			"miss error info", []byte("{\"invalid_key\":\"invalid_value\"}"),
			400, "", "{\"invalid_key\":\"invalid_value\"}", "97e2***11df", "",
		},
		{
			"invalid json", []byte("invalid json data"),
			400, "", "invalid json data", "97e2***11df", "",
		},
		{
			"null encoded message",
			[]byte("{\"error_code\":\"XXX.0001\",\"error_msg\":\"Some errors occurred.\",\"encoded_authorization_message\":null}"),
			400, "XXX.0001", "Some errors occurred.", "97e2***11df", "",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			resp := &http.Response{
				StatusCode: 400,
				Header:     mockHeader(),
				Body:       ioutil.NopCloser(bytes.NewBuffer(c.Data)),
			}
			defer resp.Body.Close()

			err := NewServiceResponseError(resp)
			assert.Equal(t, c.StatusCode, err.StatusCode)
			assert.Equal(t, c.ErrorCode, err.ErrorCode)
			assert.Equal(t, c.ErrorMessage, err.ErrorMessage)
			assert.Equal(t, c.RequestId, err.RequestId)
			assert.Equal(t, c.EncodedAuthorizationMessage, err.EncodedAuthorizationMessage)
		})
	}
}
