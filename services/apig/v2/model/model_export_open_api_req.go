package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExportOpenApiReq struct {

	// API分组发布的环境ID
	EnvId *string `json:"env_id,omitempty"`

	// API分组ID
	GroupId string `json:"group_id"`

	// 导出API的定义范围： - spec：基础定义，只包括api前端定义 - proxy：全量定义，包括api前后端定义 - all：扩展定义，包括api前后端定义以及流控、访问控制、自定义认证等扩展定义 - dev：开发定义，包括未发布的api的前后端定义
	Define *ExportOpenApiReqDefine `json:"define,omitempty"`

	// 导出的API定义的格式
	Type *ExportOpenApiReqType `json:"type,omitempty"`

	// 导出的API定义版本，默认为当前时间
	Version *string `json:"version,omitempty"`

	// 导出的API ID列表
	Apis *[]string `json:"apis,omitempty"`
}

func (o ExportOpenApiReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportOpenApiReq struct{}"
	}

	return strings.Join([]string{"ExportOpenApiReq", string(data)}, " ")
}

type ExportOpenApiReqDefine struct {
	value string
}

type ExportOpenApiReqDefineEnum struct {
	SPEC  ExportOpenApiReqDefine
	PROXY ExportOpenApiReqDefine
	ALL   ExportOpenApiReqDefine
	DEV   ExportOpenApiReqDefine
}

func GetExportOpenApiReqDefineEnum() ExportOpenApiReqDefineEnum {
	return ExportOpenApiReqDefineEnum{
		SPEC: ExportOpenApiReqDefine{
			value: "spec",
		},
		PROXY: ExportOpenApiReqDefine{
			value: "proxy",
		},
		ALL: ExportOpenApiReqDefine{
			value: "all",
		},
		DEV: ExportOpenApiReqDefine{
			value: "dev",
		},
	}
}

func (c ExportOpenApiReqDefine) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportOpenApiReqDefine) UnmarshalJSON(b []byte) error {
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

type ExportOpenApiReqType struct {
	value string
}

type ExportOpenApiReqTypeEnum struct {
	JSON ExportOpenApiReqType
	YAML ExportOpenApiReqType
	YML  ExportOpenApiReqType
}

func GetExportOpenApiReqTypeEnum() ExportOpenApiReqTypeEnum {
	return ExportOpenApiReqTypeEnum{
		JSON: ExportOpenApiReqType{
			value: "json",
		},
		YAML: ExportOpenApiReqType{
			value: "yaml",
		},
		YML: ExportOpenApiReqType{
			value: "yml",
		},
	}
}

func (c ExportOpenApiReqType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExportOpenApiReqType) UnmarshalJSON(b []byte) error {
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
