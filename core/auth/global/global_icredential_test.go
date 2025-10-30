// Copyright 2022 Huawei Technologies Co.,Ltd.
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

package global

import (
	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCredentialsBuilder_SafeBuild(t *testing.T) {
	cases := []struct {
		Name, IdpId, IdTokenFile, DomainId, ExpectedErr string
	}{
		{"IdTokenFile missed", "id", "", "domainId", "IdTokenFile is required when using IdpId&IdTokenFile"},
		{"IdpId missed", "", "file", "domainId", "IdpId is required when using IdpId&IdTokenFile"},
		{"valid params", "id", "file", "domainId", ""},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			cred, err := NewCredentialsBuilder().
				WithIdpId(c.IdpId).
				WithIdTokenFile(c.IdTokenFile).
				WithDomainId(c.DomainId).
				SafeBuild()
			if c.ExpectedErr != "" {
				assert.Error(t, err)
				var cte *sdkerr.CredentialsTypeError
				assert.True(t, errors.As(err, &cte))
				assert.Equal(t, c.ExpectedErr, cte.ErrorMessage)
			} else {
				assert.Equal(t, c.IdpId, cred.IdpId)
				assert.Equal(t, c.IdTokenFile, cred.IdTokenFile)
				assert.Equal(t, c.DomainId, cred.DomainId)
			}
		})
	}
}

func TestCredentialsBuilder_SafeBuild2(t *testing.T) {
	cases := []struct {
		Name, Ak, Sk     string
		ExpectedContains []string
	}{
		{"Empty AK", "", "sk", []string{"input ak cannot be an empty string"}},
		{"Empty SK", "ak", "", []string{"input sk cannot be an empty string"}},
		{"Empty AK&SK", "", "", []string{"input ak cannot be an empty string", "input sk cannot be an empty string"}},
	}

	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			_, err := NewCredentialsBuilder().WithAk(c.Ak).WithSk(c.Sk).SafeBuild()
			assert.IsType(t, err, &sdkerr.CredentialsTypeError{})
			msg := err.(*sdkerr.CredentialsTypeError).ErrorMessage
			for _, contain := range c.ExpectedContains {
				assert.Contains(t, msg, contain)
			}
		})
	}

	cases2 := []struct {
		Name string
		Func func() (*Credentials, error)
	}{
		{"Valid AK&SK", func() (*Credentials, error) {
			return NewCredentialsBuilder().
				WithAk("ak").
				WithSk("sk").
				SafeBuild()
		}},
		{"Invalid AK&SK then correct it", func() (*Credentials, error) {
			return NewCredentialsBuilder().
				WithAk("").
				WithSk("").
				WithAk("ak").
				WithSk("sk").
				SafeBuild()
		}},
	}

	for _, c := range cases2 {
		t.Run(c.Name, func(t *testing.T) {
			credentials, err := c.Func()
			assert.NoError(t, err)
			assert.NotEmpty(t, credentials.AK)
			assert.NotEmpty(t, credentials.SK)
		})
	}
}

func TestCredentialsBuilder_SafeBuild3(t *testing.T) {
	credentials, err := NewCredentialsBuilder().
		WithAk("ak").
		WithSk("sk").
		WithSecurityToken("token").
		WithIdTokenFile("file").
		WithIdpId("idp-id").
		WithIamEndpointOverride("iam-endpoint").
		WithDomainId("domain-id").
		SafeBuild()
	assert.NoError(t, err)
	assert.IsType(t, &Credentials{}, credentials)
	assert.Equal(t, "ak", credentials.AK)
	assert.Equal(t, "sk", credentials.SK)
	assert.Equal(t, "token", credentials.SecurityToken)
	assert.Equal(t, "idp-id", credentials.IdpId)
	assert.Equal(t, "domain-id", credentials.DomainId)
	assert.Equal(t, "iam-endpoint", credentials.IamEndpoint)
	assert.Equal(t, "file", credentials.IdTokenFile)

	credentials.AK = "new_ak"
	credentials.DomainId = "new-domain-id"
	assert.Equal(t, "new_ak", credentials.AK)
	assert.Equal(t, "new-domain-id", credentials.DomainId)
}
