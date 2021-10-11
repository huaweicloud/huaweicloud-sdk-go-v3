package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type SwaggerReq struct {
	// API分组发布的环境ID

	EnvId *string `json:"env_id,omitempty"`
	// 导出API的定义范围： - spec：基础定义，只包括api前端定义 - proxy：全量定义，包括api前后端定义 - all：扩展定义，包括api前后端定义以及流控、访问控制、自定义认证等扩展定义 - dev：开发定义，包括未发布的api的前后端定义

	Define *SwaggerReqDefine `json:"define,omitempty"`
	// 导出的API定义的格式

	Type *SwaggerReqType `json:"type,omitempty"`
	// 导出的API定义版本，默认为当前时间

	Version *string `json:"version,omitempty"`
	// 导出的API ID列表

	Apis *[]string `json:"apis,omitempty"`
	// API分组ID

	GroupId string `json:"group_id"`
}

func (o SwaggerReq) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "SwaggerReq struct{}"
	}

	return strings.Join([]string{"SwaggerReq", string(data)}, " ")
}

type SwaggerReqDefine struct {
	value string
}

type SwaggerReqDefineEnum struct {
	SPEC  SwaggerReqDefine
	PROXY SwaggerReqDefine
	ALL   SwaggerReqDefine
	DEV   SwaggerReqDefine
}

func GetSwaggerReqDefineEnum() SwaggerReqDefineEnum {
	return SwaggerReqDefineEnum{
		SPEC: SwaggerReqDefine{
			value: "spec",
		},
		PROXY: SwaggerReqDefine{
			value: "proxy",
		},
		ALL: SwaggerReqDefine{
			value: "all",
		},
		DEV: SwaggerReqDefine{
			value: "dev",
		},
	}
}

func (c SwaggerReqDefine) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *SwaggerReqDefine) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type SwaggerReqType struct {
	value string
}

type SwaggerReqTypeEnum struct {
	JSON SwaggerReqType
	YAML SwaggerReqType
	YML  SwaggerReqType
}

func GetSwaggerReqTypeEnum() SwaggerReqTypeEnum {
	return SwaggerReqTypeEnum{
		JSON: SwaggerReqType{
			value: "json",
		},
		YAML: SwaggerReqType{
			value: "yaml",
		},
		YML: SwaggerReqType{
			value: "yml",
		},
	}
}

func (c SwaggerReqType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *SwaggerReqType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
