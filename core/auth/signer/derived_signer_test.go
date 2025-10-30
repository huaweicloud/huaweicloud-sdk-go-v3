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

const (
	service  = "demo"
	regionId = "test-region-1"
)

func TestDerivedSigner_Sign(t *testing.T) {
	cases := []TestCase{
		{
			TestParam: testParam1,
			Expected: "V11-HMAC-SHA256 Credential=AccessKey/20060102/test-region-1/demo," +
				" SignedHeaders=x-sdk-date," +
				" Signature=8fd3610508884bb19718317797c7378716a94bc740d97502e569322e366341cc",
		},
		{
			TestParam: testParam2,
			Expected: "V11-HMAC-SHA256 Credential=AccessKey/20060102/test-region-1/demo, " +
				"SignedHeaders=x-sdk-date, " +
				"Signature=4112f88b92d82dd1e1b7dd142f1e99e4f1762c42ca8fe0968b3045120870a903",
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			req := buildReqWithTestcase(c)
			result, err := GetDerivedSigner().Sign(req, ak, sk, service, regionId)
			assert.NoError(t, err)
			assert.Equal(t, c.Expected, result["Authorization"])
		})
	}
}
