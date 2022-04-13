package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// Response Object
type CreateRequestPropertyResponse struct {
	// 属性ID

	PropertyId *int32 `json:"property_id,omitempty"`
	// 属性名称，首位必须为字母，支持大小写字母，数字，中划线及下划线，长度2-50

	PropertyName *string `json:"property_name,omitempty"`
	// 属性描述，长度0-200

	Description *string `json:"description,omitempty"`
	// 属性数据类型，枚举值大小写敏感；number格式为数字，范围±1.0 x 10^-28 to ±7.9228 x 10^28；sting为字符串；integer为整数；datetime为时间，格式为yyyyMMddTHHmmss；json为自定义json格式

	DataType *CreateRequestPropertyResponseDataType `json:"data_type,omitempty"`
	// 是否必填 0-非必填 1-必填

	Required *CreateRequestPropertyResponseRequired `json:"required,omitempty"`
	// 最小值，当data_type为integer或number时有效

	Min *string `json:"min,omitempty"`
	// 最大值，当data_type为integer或number时有效

	Max *string `json:"max,omitempty"`
	// 步长，当data_type为integer或number时有效

	Step *string `json:"step,omitempty"`
	// 字符串最大长度，当data_type为string, datetime, json时有效

	MaxLength *int32 `json:"max_length,omitempty"`
	// 属性单位

	Unit *string `json:"unit,omitempty"`
	// string的枚举值数组，使用逗号分隔

	EnumList       *string `json:"enum_list,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRequestPropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRequestPropertyResponse struct{}"
	}

	return strings.Join([]string{"CreateRequestPropertyResponse", string(data)}, " ")
}

type CreateRequestPropertyResponseDataType struct {
	value string
}

type CreateRequestPropertyResponseDataTypeEnum struct {
	INTEGER  CreateRequestPropertyResponseDataType
	NUMBER   CreateRequestPropertyResponseDataType
	STRING   CreateRequestPropertyResponseDataType
	DATETIME CreateRequestPropertyResponseDataType
	JSON     CreateRequestPropertyResponseDataType
}

func GetCreateRequestPropertyResponseDataTypeEnum() CreateRequestPropertyResponseDataTypeEnum {
	return CreateRequestPropertyResponseDataTypeEnum{
		INTEGER: CreateRequestPropertyResponseDataType{
			value: "integer",
		},
		NUMBER: CreateRequestPropertyResponseDataType{
			value: "number",
		},
		STRING: CreateRequestPropertyResponseDataType{
			value: "string",
		},
		DATETIME: CreateRequestPropertyResponseDataType{
			value: "datetime",
		},
		JSON: CreateRequestPropertyResponseDataType{
			value: "json",
		},
	}
}

func (c CreateRequestPropertyResponseDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRequestPropertyResponseDataType) UnmarshalJSON(b []byte) error {
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

type CreateRequestPropertyResponseRequired struct {
	value int32
}

type CreateRequestPropertyResponseRequiredEnum struct {
	E_0 CreateRequestPropertyResponseRequired
	E_1 CreateRequestPropertyResponseRequired
}

func GetCreateRequestPropertyResponseRequiredEnum() CreateRequestPropertyResponseRequiredEnum {
	return CreateRequestPropertyResponseRequiredEnum{
		E_0: CreateRequestPropertyResponseRequired{
			value: 0,
		}, E_1: CreateRequestPropertyResponseRequired{
			value: 1,
		},
	}
}

func (c CreateRequestPropertyResponseRequired) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateRequestPropertyResponseRequired) UnmarshalJSON(b []byte) error {
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
