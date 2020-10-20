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

package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/def"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/impl"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/request"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/response"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdkerr"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"reflect"
	"strings"
)

type HcHttpClient struct {
	endpoint   string
	credential auth.ICredential
	httpClient *impl.DefaultHttpClient
}

func NewHcHttpClient(httpClient *impl.DefaultHttpClient) *HcHttpClient {
	return &HcHttpClient{httpClient: httpClient}
}

func (hc *HcHttpClient) WithEndpoint(endpoint string) *HcHttpClient {
	hc.endpoint = endpoint
	return hc
}

func (hc *HcHttpClient) WithCredential(credential auth.ICredential) *HcHttpClient {
	hc.credential = credential
	return hc
}

func (hc *HcHttpClient) Sync(req interface{}, reqDef *def.HttpRequestDef, responseDef *def.HttpResponseDef) (*response.DefaultHttpResponse, error) {
	httpRequest, err := hc.buildRequest(req, reqDef)
	if err != nil {
		return nil, err
	}

	resp, err := hc.httpClient.SyncInvokeHttp(httpRequest)
	if err != nil {
		return nil, err
	}

	return hc.extractResponse(resp, responseDef)
}

func (hc *HcHttpClient) buildRequest(req interface{}, reqDef *def.HttpRequestDef) (*request.DefaultHttpRequest, error) {
	builder := request.NewHttpRequestBuilder().
		WithEndpoint(hc.endpoint).
		WithMethod(reqDef.Method).
		WithPath(reqDef.Path).
		WithBody(reqDef.BodyJson)

	if reqDef.ContentType != "" {
		builder.AddHeaderParam("Content-Type", reqDef.ContentType)
	}
	builder.AddHeaderParam("User-Agent", "huaweicloud-usdk-go/3.0")

	builder, err := hc.fillParamsFromReq(req, reqDef, builder)
	if err != nil {
		return nil, err
	}

	httpRequest, err := hc.credential.ProcessAuthRequest(builder.Build())
	if err != nil {
		return nil, err
	}
	return httpRequest, err
}

func (hc *HcHttpClient) fillParamsFromReq(req interface{}, reqDef *def.HttpRequestDef, builder *request.HttpRequestBuilder) (*request.HttpRequestBuilder, error) {
	attrMaps := hc.GetFieldJsonTags(req)

	for _, fieldDef := range reqDef.RequestFields {
		if _, ok := attrMaps[fieldDef.Name]; ok {
			value, err := hc.GetFieldValueByName(attrMaps[fieldDef.Name], req)
			if err != nil {
				return nil, err
			}
			if value == nil {
				continue
			}
			switch fieldDef.LocationType {
			case def.Header:
				builder.AddHeaderParam(fieldDef.Name, fmt.Sprintf("%v", value))
			case def.Path:
				builder.AddPathParam(fieldDef.Name, fmt.Sprintf("%v", value))
			case def.Query:
				builder.AddQueryParam(fieldDef.Name, value)
			}
		}
	}
	return builder, nil
}

func (hc *HcHttpClient) GetFieldJsonTags(structName interface{}) map[string]def.FieldJsonTag {
	attrMaps := make(map[string]def.FieldJsonTag)
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		jsonTag := t.Field(i).Tag.Get("json")
		if jsonTag != "" {
			jsonName := strings.Split(jsonTag, ",")[0]
			attrMaps[jsonName] = def.FieldJsonTag{
				FieldName: t.Field(i).Name,
				JsonName:  jsonName,
				JsonTag:   jsonTag,
			}
		}
	}
	return attrMaps
}

func (hc *HcHttpClient) GetFieldValueByName(fieldJsonTag def.FieldJsonTag, structName interface{}) (interface{}, error) {
	v := reflect.ValueOf(structName)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	value := v.FieldByName(fieldJsonTag.FieldName)
	if value.Kind() == reflect.Ptr || value.Kind() == reflect.Struct || value.Kind() == reflect.Interface {
		if value.IsNil() {
			if strings.Contains(fieldJsonTag.JsonTag, "omitempty") {
				return nil, nil
			}
			return nil, errors.New("request field " + fieldJsonTag.FieldName + " read null value")
		}
		return value.Elem(), nil
	}

	return value, nil
}

func (hc *HcHttpClient) extractResponse(resp *response.DefaultHttpResponse, responseDef *def.HttpResponseDef) (*response.DefaultHttpResponse, error) {
	if resp.GetStatusCode() >= 400 {
		return resp, sdkerr.NewServiceResponseError(resp.Response)
	}

	bodyErr := hc.deserializeResponseBody(resp, responseDef)
	if bodyErr != nil {
		return nil, bodyErr
	}

	hc.deserializeResponseHeaders(resp, responseDef)
	return resp, nil
}

func (hc *HcHttpClient) deserializeResponseBody(resp *response.DefaultHttpResponse, responseDef *def.HttpResponseDef) error {
	data, err := ioutil.ReadAll(resp.Response.Body)
	if err != nil {
		if closeErr := resp.Response.Body.Close(); closeErr != nil {
			return err
		}
		return err
	}

	if err := resp.Response.Body.Close(); err != nil {
		return err
	} else {
		resp.Response.Body = ioutil.NopCloser(bytes.NewBuffer(data))
	}

	if len(data) == 0 {
		return nil
	}

	err = jsoniter.Unmarshal(data, responseDef.BodyJson)
	if err != nil {
		if strings.HasPrefix(string(data), "{") {
			return sdkerr.ServiceResponseError{
				StatusCode:   resp.GetStatusCode(),
				RequestId:    resp.GetHeader("X-Request-Id"),
				ErrorMessage: err.Error(),
			}
		}

		dataOfListOrString := "{ \"body\" : " + string(data) + "}"
		err = jsoniter.Unmarshal([]byte(dataOfListOrString), responseDef.BodyJson)
		if err != nil {
			return sdkerr.ServiceResponseError{
				StatusCode:   resp.GetStatusCode(),
				RequestId:    resp.GetHeader("X-Request-Id"),
				ErrorMessage: err.Error(),
			}
		}
	}

	resp.BodyJson = responseDef.BodyJson
	return nil
}

func (hc *HcHttpClient) deserializeResponseHeaders(resp *response.DefaultHttpResponse, responseDef *def.HttpResponseDef) {
	headers := make(map[string]string)
	for header := range resp.Response.Header {
		headers[header] = resp.Response.Header.Get(header)
	}
	headersJsonStr, err := json.Marshal(headers)
	if err == nil && len(headersJsonStr) != 0 {
		_ = jsoniter.Unmarshal(headersJsonStr, responseDef.BodyJson)
	}
}
