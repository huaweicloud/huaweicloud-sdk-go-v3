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

package sdktime

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

type instance struct {
	Time *SdkTime `json:"time"`
}

func TestSdkTime_UnmarshalJSON(t *testing.T) {
	cases := []struct {
		Name string
		Data string
	}{
		{"format1", `{"time":"2024-05-30T14:30:00"}`},
		{"format2", `{"time":"2024-05-30T14:30:00Z"}`},
		{"format3", `{"time":"2024-05-30T14:30:00+02:00"}`},
		{"empty", `{"time":""}`},
		{"null", `{"time":null}`},
	}
	for _, c := range cases {
		t.Run(c.Name, func(t *testing.T) {
			ins := &instance{}
			err := utils.Unmarshal([]byte(c.Data), ins)
			assert.NoError(t, err)
		})
	}

}
