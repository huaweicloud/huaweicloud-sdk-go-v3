package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CreatePropertyRequestBody struct {

	// 属性名称，首位必须为字母，支持大小写字母，数字，中划线及下划线，长度2-50
	PropertyName string `json:"property_name"`

	// 属性描述，长度0-200
	Description *string `json:"description,omitempty"`

	// 属性数据类型，枚举值大小写敏感；number格式为数字，范围±1.0 x 10^-28 to ±7.9228 x 10^28；sting为字符串；integer为整数；datetime为时间，格式为yyyyMMddTHHmmss；json为自定义json格式
	DataType CreatePropertyRequestBodyDataType `json:"data_type"`

	// 是否必填 0-非必填 1-必填
	Required CreatePropertyRequestBodyRequired `json:"required"`

	// 最小值，当data_type为integer或number时必填
	Min *string `json:"min,omitempty"`

	// 最大值，当data_type为integer或number时必填
	Max *string `json:"max,omitempty"`

	// 步长，当data_type为integer或number时必填
	Step *string `json:"step,omitempty"`

	// 字符串最大长度，当data_type为string, datetime, json时必填，自动向下取整
	MaxLength *int32 `json:"max_length,omitempty"`

	// 属性单位
	Unit *string `json:"unit,omitempty"`

	// string的枚举值数组，使用逗号分隔
	EnumList *string `json:"enum_list,omitempty"`

	EnumDict *PropertyDataEnum `json:"enum_dict,omitempty"`

	// 访问模式（兼容20.0，R属性可读，W属性可写，E属性可执行）
	Method *string `json:"method,omitempty"`
}

func (o CreatePropertyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePropertyRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePropertyRequestBody", string(data)}, " ")
}

type CreatePropertyRequestBodyDataType struct {
	value string
}

type CreatePropertyRequestBodyDataTypeEnum struct {
	INTEGER  CreatePropertyRequestBodyDataType
	NUMBER   CreatePropertyRequestBodyDataType
	STRING   CreatePropertyRequestBodyDataType
	DATETIME CreatePropertyRequestBodyDataType
	JSON     CreatePropertyRequestBodyDataType
}

func GetCreatePropertyRequestBodyDataTypeEnum() CreatePropertyRequestBodyDataTypeEnum {
	return CreatePropertyRequestBodyDataTypeEnum{
		INTEGER: CreatePropertyRequestBodyDataType{
			value: "integer",
		},
		NUMBER: CreatePropertyRequestBodyDataType{
			value: "number",
		},
		STRING: CreatePropertyRequestBodyDataType{
			value: "string",
		},
		DATETIME: CreatePropertyRequestBodyDataType{
			value: "datetime",
		},
		JSON: CreatePropertyRequestBodyDataType{
			value: "json",
		},
	}
}

func (c CreatePropertyRequestBodyDataType) Value() string {
	return c.value
}

func (c CreatePropertyRequestBodyDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePropertyRequestBodyDataType) UnmarshalJSON(b []byte) error {
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

type CreatePropertyRequestBodyRequired struct {
	value int32
}

type CreatePropertyRequestBodyRequiredEnum struct {
	E_0 CreatePropertyRequestBodyRequired
	E_1 CreatePropertyRequestBodyRequired
}

func GetCreatePropertyRequestBodyRequiredEnum() CreatePropertyRequestBodyRequiredEnum {
	return CreatePropertyRequestBodyRequiredEnum{
		E_0: CreatePropertyRequestBodyRequired{
			value: 0,
		}, E_1: CreatePropertyRequestBodyRequired{
			value: 1,
		},
	}
}

func (c CreatePropertyRequestBodyRequired) Value() int32 {
	return c.value
}

func (c CreatePropertyRequestBodyRequired) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePropertyRequestBodyRequired) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(int32)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
