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

package request

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Body struct {
	Id string
}

type FileBody struct {
	File *def.FilePart `json:"file"`

	Size      *def.MultiPart `json:"size,omitempty"`
	Name      *def.MultiPart `json:"name,omitempty"`
	ShortName *def.MultiPart `json:"short_name,omitempty"`
}

func TestHttpRequestBuilder_WithBody(t *testing.T) {
	builder := NewHttpRequestBuilder().WithBody("body", &Body{Id: "id"})
	request := builder.Build()
	assert.Equal(t, &Body{Id: "id"}, request.GetBody())
}

func TestHttpRequestBuilder_WithBody2(t *testing.T) {
	tmpFile, err := os.CreateTemp("", "go-test-*.tmp")
	assert.NoError(t, err)
	defer func() {
		err = tmpFile.Close()
		assert.NoError(t, err)
		err = os.Remove(tmpFile.Name())
		assert.NoError(t, err)
	}()

	_, err = tmpFile.Write([]byte{1, 2, 3})
	assert.NoError(t, err)

	body := &FileBody{
		File: def.NewFilePart(tmpFile),
		Size: def.NewMultiPart(3),
		Name: def.NewMultiPart("tmp"),
	}
	request := NewHttpRequestBuilder().WithBody("multipart", body).Build()
	assert.NotEmpty(t, request.GetFormParams())
	assert.Equal(t, 3, len(request.GetFormParams()))
}
