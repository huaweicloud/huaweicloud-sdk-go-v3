// Copyright 2020 Huawei Technologies Co.,Ltd.
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
	"bytes"
	"crypto/tls"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/config"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/httphandler"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/response"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type DefaultHttpClient struct {
	httpHandler  *httphandler.HttpHandler
	httpConfig   *config.HttpConfig
	transport    *http.Transport
	goHttpClient *http.Client
}

func NewDefaultHttpClient(httpConfig *config.HttpConfig) *DefaultHttpClient {
	transport := httpConfig.HttpTransport
	if transport == nil {
		transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: httpConfig.IgnoreSSLVerification},
		}
	}

	if httpConfig.DialContext != nil {
		transport.DialContext = httpConfig.DialContext
	}

	if httpConfig.HttpProxy != nil {
		proxyUrl := httpConfig.HttpProxy.GetProxyUrl()
		if proxyUrl != "" {
			proxy, _ := url.Parse(proxyUrl)
			transport.Proxy = http.ProxyURL(proxy)
		}
	}

	client := &DefaultHttpClient{
		transport:  transport,
		httpConfig: httpConfig,
	}

	client.goHttpClient = &http.Client{
		Transport: client.transport,
		Timeout:   httpConfig.Timeout,
	}

	client.httpHandler = httpConfig.HttpHandler

	return client
}

func (client *DefaultHttpClient) SyncInvokeHttp(request *request.DefaultHttpRequest) (*response.DefaultHttpResponse, error) {
	req, err := request.ConvertRequest()
	if err != nil {
		return nil, err
	}

	if client.httpHandler != nil && client.httpHandler.RequestHandlers != nil && req != nil {
		bodyBytes, err := httputil.DumpRequest(req, true)
		if err != nil {
			return nil, err
		}
		reqClone := req.Clone(req.Context())
		reqClone.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		defer reqClone.Body.Close()
		client.httpHandler.RequestHandlers(*reqClone)
	}

	startTime := time.Now()

	var resp *http.Response
	tried := 0
	for {
		resp, err = client.goHttpClient.Do(req)
		tried++
		if client.httpConfig.Retries <= 1 || tried >= client.httpConfig.Retries ||
			err == nil || (resp != nil && resp.StatusCode < 300) {
			break
		}
	}

	endTime := time.Now()

	if client.httpHandler != nil && client.httpHandler.ResponseHandlers != nil && resp != nil {
		bodyBytes, err := httputil.DumpResponse(resp, true)
		if err != nil {
			return nil, err
		}
		respClone := http.Response{
			Body:             ioutil.NopCloser(bytes.NewBuffer(bodyBytes)),
			Status:           resp.Status,
			StatusCode:       resp.StatusCode,
			Proto:            resp.Proto,
			ProtoMajor:       resp.ProtoMajor,
			ProtoMinor:       resp.ProtoMinor,
			Header:           resp.Header,
			ContentLength:    resp.ContentLength,
			TransferEncoding: resp.TransferEncoding,
			Close:            resp.Close,
			Uncompressed:     resp.Uncompressed,
			Trailer:          resp.Trailer,
		}
		defer respClone.Body.Close()
		client.httpHandler.ResponseHandlers(respClone)
	}

	if client.httpHandler != nil && client.httpHandler.MonitorHandlers != nil {
		metric := &httphandler.MonitorMetric{
			Host:      req.URL.Host,
			Method:    req.Method,
			Path:      req.URL.Path,
			Raw:       req.URL.RawQuery,
			UserAgent: req.UserAgent(),
			Latency:   endTime.Sub(startTime),
		}

		if resp != nil {
			metric.RequestId = resp.Header.Get("X-Request-Id")
			metric.StatusCode = resp.StatusCode
			metric.ContentLength = resp.ContentLength
		}

		client.httpHandler.MonitorHandlers(metric)
	}

	if err != nil {
		return nil, err
	}

	return response.NewDefaultHttpResponse(resp), nil
}
