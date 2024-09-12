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
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/signer/algorithm"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type MockRoundTripper struct {
}

func (r *MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return nil, errors.New("test")
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

func TestNewDefaultHttpClient(t *testing.T) {
	httpConfig := config.DefaultHttpConfig()
	httpConfig.WithIgnoreSSLVerification(true)
	httpConfig.WithTimeout(120)
	proxy := config.NewProxy().
		// 请根据实际情况替换示例中的代理协议、地址和端口号
		WithSchema("http").
		WithHost("proxy.test.com").
		WithPort(80).
		// 如果代理需要认证，请配置用户名和密码
		WithUsername("user").
		WithPassword("pass")
	httpConfig.WithProxy(proxy)
	dialContext := func(ctx context.Context, network string, addr string) (net.Conn, error) {
		return nil, errors.New("test")
	}
	httpConfig.WithDialContext(dialContext)
	httpConfig.WithHttpTransport(&http.Transport{})
	requestHandler := func(request http.Request) {}
	responseHandler := func(response http.Response) {}
	httpHandler := httphandler.NewHttpHandler().AddRequestHandler(requestHandler).AddResponseHandler(responseHandler)
	httpConfig.WithHttpHandler(httpHandler)

	client := NewDefaultHttpClient(httpConfig)
	assert.NotNil(t, client)
	assert.Equal(t, httpConfig, client.httpConfig)
}

func TestNewDefaultHttpClient_TransPort(t *testing.T) {
	conf := config.DefaultHttpConfig()
	client := NewDefaultHttpClient(conf)
	assert.IsType(t, &http.Transport{}, client.roundTripper)

	conf1 := config.DefaultHttpConfig()
	conf1.WithHttpTransport(&http.Transport{})
	client1 := NewDefaultHttpClient(conf1)
	assert.IsType(t, &http.Transport{}, client1.roundTripper)

	conf2 := config.DefaultHttpConfig()
	conf2.WithHttpRoundTripper(&MockRoundTripper{})
	client2 := NewDefaultHttpClient(conf2)
	assert.IsType(t, &MockRoundTripper{}, client2.roundTripper)

	conf3 := config.DefaultHttpConfig()
	conf3.WithHttpTransport(&http.Transport{})
	conf3.WithHttpRoundTripper(&MockRoundTripper{})
	client3 := NewDefaultHttpClient(conf3)
	assert.IsType(t, &http.Transport{}, client3.roundTripper)
}
