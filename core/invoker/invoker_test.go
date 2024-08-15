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

package invoker

import (
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/basic"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/invoker/retry"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockModel struct {
	HttpStatusCode int `json:"-"`
}

var reqDef = def.NewHttpRequestDefBuilder().
	WithMethod("GET").WithPath("/").
	WithContentType("application/json").
	WithResponse(&mockModel{}).
	Build()

func mockHcClient(url string, options ...interface{}) (*core.HcHttpClient, error) {
	var credentials auth.ICredential
	credentials, err := basic.NewCredentialsBuilder().WithAk("ak").WithSk("sk").SafeBuild()
	if err != nil {
		return nil, err
	}
	conf := config.DefaultHttpConfig()

	for _, op := range options {
		switch op.(type) {
		case auth.ICredential:
			credentials = op.(auth.ICredential)
		case *config.HttpConfig:
			conf = op.(*config.HttpConfig)
		default:
			return nil, errors.New("invalid option type")
		}
	}

	client, err := core.NewHcHttpClientBuilder().
		WithHttpConfig(conf).
		WithCredential(credentials).
		WithEndpoints([]string{url}).
		SafeBuild()
	if err != nil {
		return nil, err
	}

	return client, nil
}

func TestBaseInvoker_Exchange_Attributes(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, "{}")
		assert.Nil(t, err)
	}))
	defer ts.Close()

	handler := httphandler.NewHttpHandler().AddMonitorHandler(func(metric *httphandler.MonitorMetric) {
		assert.Equal(t, map[string]interface{}{"key": "value"}, metric.Attributes)
	})
	httpConfig := config.DefaultHttpConfig().WithHttpHandler(handler)

	client, err := mockHcClient(ts.URL, httpConfig)
	assert.Nil(t, err)
	invoker := NewBaseInvoker(client, &mockModel{}, reqDef).WithRetry(10, func(resp interface{}, err error) bool {
		return err != nil
	}, retry.NewDecorRelatedJitter())
	invoker.Exchange.Attributes = map[string]interface{}{"key": "value"}
	resp, err := invoker.Invoke()
	assert.Nil(t, err)
	assert.IsType(t, &mockModel{}, resp)
	assert.Equal(t, 200, resp.(*mockModel).HttpStatusCode)
}

// 当发生ServerResponseException(status_code>=500)时进行重试,最大重试次数设置为3
// 预期状态码为502,调用次数为4(正常调用1次+请求重试3次)
func TestBaseInvoker_WithRetry(t *testing.T) {
	count := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, err := fmt.Fprintln(w, "{}")
		count++
		assert.Nil(t, err)
	}))
	defer ts.Close()

	client, err := mockHcClient(ts.URL)
	assert.Nil(t, err)
	invoker := NewBaseInvoker(client, &mockModel{}, reqDef).WithRetry(3, func(resp interface{}, err error) bool {
		return err != nil
	}, &retry.None{})
	resp, err := invoker.Invoke()
	assert.Nil(t, resp)
	assert.IsType(t, &sdkerr.ServiceResponseError{}, err)
	assert.Equal(t, 502, err.(*sdkerr.ServiceResponseError).StatusCode)
	assert.Equal(t, 4, count)
}

// 当发生ServerResponseException(status_code>=500)时进行重试,最大重试次数设置为10,第一次重试成功
// 预期状态码为200,调用次数为2(正常调用1次+请求重试1次)
func TestBaseInvoker_WithRetry2(t *testing.T) {
	count := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if count == 0 {
			w.WriteHeader(http.StatusBadGateway)
		}
		_, err := fmt.Fprintln(w, "{}")
		count++
		assert.Nil(t, err)
	}))
	defer ts.Close()

	client, err := mockHcClient(ts.URL)
	assert.Nil(t, err)
	invoker := NewBaseInvoker(client, &mockModel{}, reqDef).WithRetry(10, func(resp interface{}, err error) bool {
		return err != nil
	}, retry.NewDecorRelatedJitter())
	resp, err := invoker.Invoke()
	assert.Nil(t, err)
	assert.IsType(t, &mockModel{}, resp)
	assert.Equal(t, 200, resp.(*mockModel).HttpStatusCode)
	assert.Equal(t, 2, count)
}

// 在刚好达到最大重试次数时请求成功
// 预期状态码为200,调用次数为3(正常调用1次+请求重试2次)
func TestBaseInvoker_WithRetry3(t *testing.T) {
	count := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if count < 2 {
			w.WriteHeader(http.StatusBadGateway)
		}
		_, err := fmt.Fprintln(w, "{}")
		count++
		assert.Nil(t, err)
	}))
	defer ts.Close()

	client, err := mockHcClient(ts.URL)
	assert.Nil(t, err)
	invoker := NewBaseInvoker(client, &mockModel{}, reqDef).WithRetry(2, func(resp interface{}, err error) bool {
		return err != nil
	}, retry.NewEqualJitter())
	resp, err := invoker.Invoke()
	assert.Nil(t, err)
	assert.IsType(t, &mockModel{}, resp)
	assert.Equal(t, 200, resp.(*mockModel).HttpStatusCode)
	assert.Equal(t, 3, count)
}

// 重试一次,预期调用次数为2 (1次正常调用+1次重试)
func TestBaseInvoker_WithRetry4(t *testing.T) {
	count := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, err := fmt.Fprintln(w, "{}")
		count++
		assert.Nil(t, err)
	}))
	defer ts.Close()

	client, err := mockHcClient(ts.URL)
	assert.Nil(t, err)
	invoker := NewBaseInvoker(client, &mockModel{}, reqDef).WithRetry(1, func(resp interface{}, err error) bool {
		return err != nil
	}, retry.NewRandomJitter())
	resp, err := invoker.Invoke()
	assert.Nil(t, resp)
	assert.IsType(t, &sdkerr.ServiceResponseError{}, err)
	assert.Equal(t, 502, err.(*sdkerr.ServiceResponseError).StatusCode)
	assert.Equal(t, 2, count)
}

// 不满足重试条件不重试,预期调用次数为1 (1次正常调用+0次重试)
func TestBaseInvoker_WithRetry5(t *testing.T) {
	count := 0
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadGateway)
		_, err := fmt.Fprintln(w, "{}")
		count++
		assert.Nil(t, err)
	}))
	defer ts.Close()

	client, err := mockHcClient(ts.URL)
	assert.Nil(t, err)
	invoker := NewBaseInvoker(client, &mockModel{}, reqDef).WithRetry(10, func(resp interface{}, err error) bool {
		return false
	}, retry.None{})
	resp, err := invoker.Invoke()
	assert.Nil(t, resp)
	assert.IsType(t, &sdkerr.ServiceResponseError{}, err)
	assert.Equal(t, 502, err.(*sdkerr.ServiceResponseError).StatusCode)
	assert.Equal(t, 1, count)
}
