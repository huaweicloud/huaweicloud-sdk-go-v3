package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreateEdgeApplicationRequestDto struct {

	// 应用ID
	EdgeAppId string `json:"edge_app_id"`

	// 应用名称
	EdgeAppName *string `json:"edge_app_name,omitempty"`

	// 应用描述
	Description *string `json:"description,omitempty"`

	// 功能类型,分为数据处理（DATA_PROCESSING）、协议解析（PROTOCOL_PARSING）、IT集成（ON_PREMISE_INTEGRATION）和混合应用（COMPOSITE_APPLICATION），数据默认为DATA_PROCESSING，数据处理模块可以传输消息，协议解析为驱动类型，IT集成为部署南向3rdIA使用，混合应用可同时实现数据处理和协议解析。
	FunctionType *CreateEdgeApplicationRequestDtoFunctionType `json:"function_type,omitempty"`

	// 驱动协议类型OPCUA|Modbus-TCP
	Protocol *string `json:"protocol,omitempty"`
}

func (o CreateEdgeApplicationRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateEdgeApplicationRequestDto struct{}"
	}

	return strings.Join([]string{"CreateEdgeApplicationRequestDto", string(data)}, " ")
}

type CreateEdgeApplicationRequestDtoFunctionType struct {
	value string
}

type CreateEdgeApplicationRequestDtoFunctionTypeEnum struct {
	DATA_PROCESSING        CreateEdgeApplicationRequestDtoFunctionType
	PROTOCOL_PARSING       CreateEdgeApplicationRequestDtoFunctionType
	ON_PREMISE_INTEGRATION CreateEdgeApplicationRequestDtoFunctionType
	GATEWAY_MANAGER        CreateEdgeApplicationRequestDtoFunctionType
	COMPOSITE_APPLICATION  CreateEdgeApplicationRequestDtoFunctionType
	DATA_COLLECTION        CreateEdgeApplicationRequestDtoFunctionType
}

func GetCreateEdgeApplicationRequestDtoFunctionTypeEnum() CreateEdgeApplicationRequestDtoFunctionTypeEnum {
	return CreateEdgeApplicationRequestDtoFunctionTypeEnum{
		DATA_PROCESSING: CreateEdgeApplicationRequestDtoFunctionType{
			value: "DATA_PROCESSING",
		},
		PROTOCOL_PARSING: CreateEdgeApplicationRequestDtoFunctionType{
			value: "PROTOCOL_PARSING",
		},
		ON_PREMISE_INTEGRATION: CreateEdgeApplicationRequestDtoFunctionType{
			value: "ON_PREMISE_INTEGRATION",
		},
		GATEWAY_MANAGER: CreateEdgeApplicationRequestDtoFunctionType{
			value: "GATEWAY_MANAGER",
		},
		COMPOSITE_APPLICATION: CreateEdgeApplicationRequestDtoFunctionType{
			value: "COMPOSITE_APPLICATION",
		},
		DATA_COLLECTION: CreateEdgeApplicationRequestDtoFunctionType{
			value: "DATA_COLLECTION",
		},
	}
}

func (c CreateEdgeApplicationRequestDtoFunctionType) Value() string {
	return c.value
}

func (c CreateEdgeApplicationRequestDtoFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateEdgeApplicationRequestDtoFunctionType) UnmarshalJSON(b []byte) error {
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
