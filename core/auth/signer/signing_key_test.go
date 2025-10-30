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

func TestP256SigningKey_Sign(t *testing.T) {
	data := []byte("HelloWorld")
	signingKey, err := p256sha256SignerInst.GetSigningKey(ak, sk)
	assert.NoError(t, err)
	_, ok := signingKey.(P256SigningKey)
	assert.True(t, ok)

	sig, err := signingKey.Sign(data)
	assert.NoError(t, err)
	assert.True(t, signingKey.Verify(sig, data))
}

func TestSM2SigningKey_Sign(t *testing.T) {
	data := []byte("HelloWorld")
	signingKey, err := sm2sm3SignerInst.GetSigningKey(ak, sk)
	assert.NoError(t, err)
	_, ok := signingKey.(SM2SigningKey)
	assert.True(t, ok)

	sig, err := signingKey.Sign(data)
	assert.NoError(t, err)
	assert.True(t, signingKey.Verify(sig, data))
}
