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
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSM3Signer_Sign(t *testing.T) {
	cases := []TestCase{
		{
			TestParam: testParam1,
			Expected: "SDK-HMAC-SM3 Access=AccessKey, SignedHeaders=x-sdk-date, " +
				"Signature=acdeecf8061419275127135b54532c2c20b683bf9bbb8a32bee021a8c40befd4",
		},
		{
			TestParam: testParam2,
			Expected: "SDK-HMAC-SM3 Access=AccessKey, SignedHeaders=x-sdk-date, " +
				"Signature=359e49285fe5a469f82b86fd761b1a4a7ee16e3181de7d643c53a30667cde508",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			req := buildReqWithTestcase(c)
			result, err := sm3SignerInst.Sign(req, ak, sk)
			assert.Nil(t, err)
			assert.Equal(t, c.Expected, result["Authorization"])
		})
	}
}
