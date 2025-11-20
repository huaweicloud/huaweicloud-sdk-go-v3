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
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/signer/algorithm"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/progress"
	"github.com/stretchr/testify/assert"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

type RequestBody struct {
	Id string `json:"id"`
}

type Listener struct {
}

func (l Listener) ProgressChanged(event *progress.Event) {
}

func TestDefaultHttpRequest_ConvertRequest(t *testing.T) {
	// application/json
	q1 := "q1"
	q2 := []string{"a", "b"}
	q3 := map[string]string{"key": "val"}

	req := NewHttpRequestBuilder().
		WithMethod(http.MethodPost).
		WithPath("v1/{project_id}/{app_id}").
		WithEndpoint("endpoint").
		WithSigningAlgorithm(algorithm.GetDefaultSigningAlgorithm()).
		WithBody("Body", &RequestBody{Id: "id_value"}).
		AddAutoFilledPathParam("project_id", "project").
		AddPathParam("app_id", "app").
		AddHeaderParam("Content-Type", "application/json").
		AddQueryParam("q1", reflect.ValueOf(q1)).
		AddQueryParam("q2", reflect.ValueOf(q2)).
		AddQueryParam("q3", reflect.ValueOf(q3)).
		Build()

	httpReq, err := req.ConvertRequest()
	assert.NoError(t, err)
	fmt.Println(httpReq.URL)

	// application/x-www-form-urlencoded
	req = NewHttpRequestBuilder().
		WithMethod(http.MethodPost).
		WithEndpoint("endpoint").
		AddFormParam("name", def.NewMultiPart("value")).
		AddFormParam("count", def.NewMultiPart(1)).
		AddHeaderParam("Content-Type", "application/x-www-form-urlencoded").
		Build()

	httpReq, err = req.ConvertRequest()
	assert.NoError(t, err)
	defer func(Body io.ReadCloser) {
		assert.NoError(t, Body.Close())
	}(httpReq.Body)
	b, err := ioutil.ReadAll(httpReq.Body)
	assert.Equal(t, "count=1&name=value", string(b))

	// multipart/form-data
	getwd, err := os.Getwd()
	assert.NoError(t, err)
	file := filepath.Join(getwd, "test.txt")
	open, err := os.Open(file)
	assert.NoError(t, err)

	req = NewHttpRequestBuilder().
		WithMethod(http.MethodPost).
		WithEndpoint("endpoint").
		AddFormParam("file", def.NewFilePart(open)).
		AddFormParam("name", def.NewMultiPart("value")).
		AddFormParam("count", def.NewMultiPart(123456789)).
		AddHeaderParam("Content-Type", "multipart/form-data").
		Build()

	httpReq, err = req.ConvertRequest()
	assert.NoError(t, err)
	assert.NoError(t, open.Close())
	b, err = ioutil.ReadAll(httpReq.Body)
	assert.NoError(t, err)
	assert.NoError(t, httpReq.Body.Close())
	content := string(b)
	assert.Contains(t, content, `"name"`)
	assert.Contains(t, content, "value")
	assert.Contains(t, content, `"count"`)
	assert.Contains(t, content, "123456789")
	assert.Contains(t, content, `"file"`)
	assert.Contains(t, content, "test-content")

	// octet-stream
	open, err = os.Open(file)
	assert.NoError(t, err)

	req = NewHttpRequestBuilder().
		WithMethod(http.MethodPost).
		WithEndpoint("endpoint").
		WithBody("File", *open).
		AddHeaderParam("Content-Type", "application/octet-stream").
		WithProgressInterval(100).
		WithProgressListener(Listener{}).
		Build()
	httpReq, err = req.ConvertRequest()
	assert.NoError(t, err)
	b, err = ioutil.ReadAll(httpReq.Body)
	assert.NoError(t, err)
	assert.NoError(t, httpReq.Body.Close())
	content = strings.TrimSpace(string(b))
	assert.Equal(t, "test-content", content)
}

func TestDefaultHttpRequest_GetBodyToBytes(t *testing.T) {
	req := NewHttpRequestBuilder().
		WithBody("Body", &RequestBody{Id: "id_value"}).
		AddHeaderParam("Content-Type", "application/xml").
		Build()
	b, err := req.GetBodyToBytes()
	assert.NoError(t, err)
	assert.Equal(t, "<RequestBody><Id>id_value</Id></RequestBody>", b.String())

	req = NewHttpRequestBuilder().
		WithBody("Body", &RequestBody{Id: "id_value"}).
		AddHeaderParam("Content-Type", "application/bson").
		Build()
	b, err = req.GetBodyToBytes()
	assert.NoError(t, err)
	assert.Equal(t, []byte{0x16, 0x0, 0x0, 0x0, 0x2, 0x69, 0x64, 0x0, 0x9, 0x0, 0x0, 0x0,
		0x69, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x0, 0x0}, b.Bytes())
}
