package model

import (
	"encoding/json"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Request Object
type ExportFunctionRequest struct {
	// 函数的URN（Uniform Resource Name），唯一标识函数。

	FunctionUrn string `json:"function_urn"`
	// 是否导出函数配置

	Config *bool `json:"config,omitempty"`
	// 是否导出函数代码

	Code *bool `json:"code,omitempty"`
	// 兼容老的方式，type=code代表导出代码,type=config代码导出配置

	Type *ExportFunctionRequestType `json:"type,omitempty"`
}

func (o ExportFunctionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ExportFunctionRequest struct{}"
	}

	return strings.Join([]string{"ExportFunctionRequest", string(data)}, " ")
}

type ExportFunctionRequestType struct {
	value string
}

type ExportFunctionRequestTypeEnum struct {
	TYPE ExportFunctionRequestType
	CODE ExportFunctionRequestType
}

func GetExportFunctionRequestTypeEnum() ExportFunctionRequestTypeEnum {
	return ExportFunctionRequestTypeEnum{
		TYPE: ExportFunctionRequestType{
			value: "type",
		},
		CODE: ExportFunctionRequestType{
			value: "code",
		},
	}
}

func (c ExportFunctionRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.value)
}

func (c *ExportFunctionRequestType) UnmarshalJSON(b []byte) error {
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
