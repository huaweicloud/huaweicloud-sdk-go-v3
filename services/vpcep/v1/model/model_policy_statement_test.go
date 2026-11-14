// Copyright 2026 Huawei Technologies Co.,Ltd.
//
// Licensed under the Apache Software Foundation (ASF) under one
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

package model

import (
	"testing"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/stretchr/testify/assert"
)

// Issue #111: gateway endpoints (OBS/SFS) return policy_statement with
// "Action" as a plain string instead of an array of strings. A page
// containing such an endpoint must still decode successfully.
func TestUnmarshalListEndpointsResponseWithStringAction(t *testing.T) {
	body := `{
		"endpoints": [
			{
				"id": "4189d3c2-8882-4871-a3c2-d380272eed88",
				"service_type": "gateway",
				"endpoint_service_name": "com.mycompany.obs.cn-north-4.test",
				"policy_statement": [
					{"Effect": "Allow", "Action": "*", "Resource": ["*", "*/*"]}
				]
			},
			{
				"id": "82b299ae-ea2c-498c-a7bc-9f1ad4a91b41",
				"service_type": "interface",
				"policy_statement": [
					{"Effect": "Allow", "Action": ["obs:*"], "Resource": ["*"]}
				]
			}
		],
		"total_count": 2
	}`

	var resp ListEndpointsResponse
	err := utils.Unmarshal([]byte(body), &resp)
	if !assert.NoError(t, err) {
		return
	}
	if !assert.NotNil(t, resp.Endpoints) || !assert.Len(t, *resp.Endpoints, 2) {
		return
	}

	gatewayStatements := (*resp.Endpoints)[0].PolicyStatement
	if !assert.NotNil(t, gatewayStatements) || !assert.Len(t, *gatewayStatements, 1) {
		return
	}
	assert.Equal(t, "Allow", (*gatewayStatements)[0].Effect.Value())
	assert.Equal(t, []string{"*"}, (*gatewayStatements)[0].Action)
	assert.Equal(t, []string{"*", "*/*"}, (*gatewayStatements)[0].Resource)

	interfaceStatements := (*resp.Endpoints)[1].PolicyStatement
	if !assert.NotNil(t, interfaceStatements) || !assert.Len(t, *interfaceStatements, 1) {
		return
	}
	assert.Equal(t, []string{"obs:*"}, (*interfaceStatements)[0].Action)
	assert.Equal(t, []string{"*"}, (*interfaceStatements)[0].Resource)
}

// The VPCEP service may also return "Resource" as a plain string.
func TestUnmarshalPolicyStatementWithStringResource(t *testing.T) {
	var statement PolicyStatement
	err := utils.Unmarshal([]byte(`{"Effect": "Allow", "Action": ["obs:*"], "Resource": "*"}`), &statement)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "Allow", statement.Effect.Value())
	assert.Equal(t, []string{"obs:*"}, statement.Action)
	assert.Equal(t, []string{"*"}, statement.Resource)
}

// Both "Action" and "Resource" may be plain strings at the same time.
func TestUnmarshalPolicyStatementWithStringActionAndResource(t *testing.T) {
	var statement PolicyStatement
	err := utils.Unmarshal([]byte(`{"Effect": "Allow", "Action": "*", "Resource": "*"}`), &statement)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "Allow", statement.Effect.Value())
	assert.Equal(t, []string{"*"}, statement.Action)
	assert.Equal(t, []string{"*"}, statement.Resource)
}

// Guards existing behavior: the documented array form and the other fields
// keep decoding as before.
func TestUnmarshalPolicyStatementWithSliceActionAndResource(t *testing.T) {
	var statement PolicyStatement
	err := utils.Unmarshal([]byte(`{
		"Effect": "Allow",
		"Action": ["obs:bucket:GetObject"],
		"Resource": ["obs:*:*:bucket:*"],
		"Condition": {"ip_address": ["10.0.0.0/8"]},
		"Sid": "sid-01"
	}`), &statement)
	if !assert.NoError(t, err) {
		return
	}

	assert.Equal(t, "Allow", statement.Effect.Value())
	assert.Equal(t, []string{"obs:bucket:GetObject"}, statement.Action)
	assert.Equal(t, []string{"obs:*:*:bucket:*"}, statement.Resource)
	assert.NotNil(t, statement.Condition)
	if assert.NotNil(t, statement.Sid) {
		assert.Equal(t, "sid-01", *statement.Sid)
	}
}

// Guards existing behavior: missing or null Action/Resource decode to nil.
func TestUnmarshalPolicyStatementWithMissingOrNullFields(t *testing.T) {
	var missing PolicyStatement
	err := utils.Unmarshal([]byte(`{"Effect": "Allow"}`), &missing)
	if !assert.NoError(t, err) {
		return
	}
	assert.Nil(t, missing.Action)
	assert.Nil(t, missing.Resource)

	var nullFields PolicyStatement
	err = utils.Unmarshal([]byte(`{"Effect": "Allow", "Action": null, "Resource": null}`), &nullFields)
	assert.NoError(t, err)
	assert.Nil(t, nullFields.Action)
	assert.Nil(t, nullFields.Resource)
}

// Guards the request direction: marshaling must keep producing arrays,
// as required by the documented request format.
func TestMarshalPolicyStatementProducesSlices(t *testing.T) {
	statement := PolicyStatement{
		Effect:   GetPolicyStatementEffectEnum().ALLOW,
		Action:   []string{"obs:*"},
		Resource: []string{"*"},
	}

	data, err := utils.Marshal(statement)
	if !assert.NoError(t, err) {
		return
	}
	assert.JSONEq(t, `{"Effect":"Allow","Action":["obs:*"],"Resource":["*"]}`, string(data))
}
