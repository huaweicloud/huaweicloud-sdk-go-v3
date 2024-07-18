// Copyright 2024 Huawei Technologies Co.,Ltd.
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

package impl

import (
	"encoding/json"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/signer/algorithm"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestDefaultHttpClient_SyncInvokeHttp(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method != "POST" {
			http.Error(w, fmt.Sprintf("{\"code\": 400, \"message\": \"Invalid request method: %s\"}", r.Method), http.StatusBadRequest)
			return
		}

		defer r.Body.Close()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, fmt.Sprintf("{\"code\": 400, \"message\": \"Error reading request body: %s\"}", err.Error()), http.StatusBadRequest)
			return
		}

		var person Person
		if err := json.Unmarshal(body, &person); err != nil {
			http.Error(w, fmt.Sprintf("{\"code\": 400, \"message\": \"Error parsing JSON: %s\"}", err.Error()), http.StatusBadRequest)
			return
		}

		_, err = fmt.Fprintf(w, "{\"name\": \"%s\", \"age\": %d}", person.Name, person.Age)
		assert.Nil(t, err)
	}))
	defer ts.Close()

	handler := httphandler.NewHttpHandler().AddRequestHandler(func(req http.Request) {
		defer req.Body.Close()
		bytes, err := ioutil.ReadAll(req.Body)
		assert.Nil(t, err)

		var person Person
		err = json.Unmarshal(bytes, &person)
		assert.Nil(t, err)
		assert.Equal(t, "Miles", person.Name)
		assert.Equal(t, 18, person.Age)
	}).AddResponseHandler(func(resp http.Response) {
		assert.Equal(t, 200, resp.StatusCode)
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		assert.Nil(t, err)

		var person Person
		err = json.Unmarshal(bytes, &person)
		assert.Nil(t, err)
		assert.Equal(t, "Miles", person.Name)
		assert.Equal(t, 18, person.Age)
	}).AddMonitorHandler(func(metric *httphandler.MonitorMetric) {
		assert.Equal(t, "POST", metric.Method)
		assert.Equal(t, "/", metric.Path)
		assert.Equal(t, 200, metric.StatusCode)
		assert.Equal(t, int64(28), metric.ContentLength)
	})
	httpConfig := config.DefaultHttpConfig().WithIgnoreSSLVerification(true).WithHttpHandler(handler)

	client := NewDefaultHttpClient(httpConfig)
	httpRequest := request.NewHttpRequestBuilder().
		WithMethod("POST").
		WithEndpoint(ts.URL).
		WithPath("/").
		WithSigningAlgorithm(algorithm.HmacSHA256).
		WithBody("Person", &Person{
			Name: "Miles",
			Age:  18,
		}).Build()
	httpResponse, err := client.SyncInvokeHttp(httpRequest)
	assert.Nil(t, err)
	assert.Equal(t, 200, httpResponse.GetStatusCode())
	body, err := httpResponse.GetBodyAsString()
	assert.Nil(t, err)
	assert.Equal(t, "{\"name\": \"Miles\", \"age\": 18}", body)
}
