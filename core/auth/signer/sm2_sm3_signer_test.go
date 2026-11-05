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
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSM2SM3Signer_Sign(t *testing.T) {
	expectedPrefix := "SDK-SM2-SM3"
	cases := []TestCase{
		{
			TestParam: testParam1,
			Expected:  expectedPrefix,
		},
		{
			TestParam: testParam2,
			Expected:  expectedPrefix,
		},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			req := buildReqWithTestcase(c)
			result, err := sm2sm3SignerInst.Sign(req, ak, sk)
			assert.NoError(t, err)
			assert.Contains(t, result["Authorization"], c.Expected)
		})
	}
}

func TestSM2SM3Signer_RoundTrip(t *testing.T) {
	for _, tp := range []TestParam{testParam1, testParam2} {
		req := buildReqWithTestcase(TestCase{TestParam: tp})

		result, err := sm2sm3SignerInst.Sign(req, ak, sk)
		assert.NoError(t, err)

		authHeader := result["Authorization"]
		sigIdx := strings.Index(authHeader, "Signature=")
		assert.NotEqual(t, -1, sigIdx)
		sigHex := authHeader[sigIdx+len("Signature="):]
		sigBytes, err := hex.DecodeString(sigHex)
		assert.NoError(t, err)

		originalHeaders := req.GetHeaderParams()
		timeVal := extractTime(originalHeaders)
		signedHeaders := extractSignedHeaders(originalHeaders)
		cr, err := canonicalRequest(req, signedHeaders, xSdkContentSm3, sm3HasherInst)
		assert.NoError(t, err)
		sts, err := stringToSign(sdkSm2Sm3, cr, timeVal, sm3HasherInst)
		assert.NoError(t, err)

		signingKey, err := sm2sm3SignerInst.GetSigningKey(ak, sk)
		assert.NoError(t, err)

		ok := signingKey.Verify(sigBytes, []byte(sts))
		assert.True(t, ok, "sign->verify round-trip should succeed")
	}
}
