package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreatePropertyResponse Response Object
type CreatePropertyResponse struct {

	// 属性ID
	PropertyId *int32 `json:"property_id,omitempty"`

	// 属性名称，首位必须为字母，支持大小写字母，数字，中划线及下划线，长度2-50
	PropertyName *string `json:"property_name,omitempty"`

	// 属性描述，长度0-200
	Description *string `json:"description,omitempty"`

	// 属性数据类型，boolean枚举值大小写敏感；number格式为数字，范围±1.0 x 10^-28 to ±7.9228 x 10^28；string为字符串；integer为整数；datetime为时间，格式为yyyyMMddTHHmmss；json为自定义json格式; array为数组类型
	DataType *CreatePropertyResponseDataType `json:"data_type,omitempty"`

	// 是否必填 0-非必填 1-必填
	Required *CreatePropertyResponseRequired `json:"required,omitempty"`

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
	EnumList *string `json:"enum_list,omitempty"`

	// 当数据类型为boolean枚举值时填写json格式数据，形如\"enum_dict\":{\"0\":\"xxx\",\"1\":\"xxx\"}
	EnumDict *interface{} `json:"enum_dict,omitempty"`

	// 访问模式（兼容20.0，R属性可读，W属性可写，E属性可执行）
	Method         *string `json:"method,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePropertyResponse struct{}"
	}

	return strings.Join([]string{"CreatePropertyResponse", string(data)}, " ")
}

type CreatePropertyResponseDataType struct {
	value string
}

type CreatePropertyResponseDataTypeEnum struct {
	INTEGER  CreatePropertyResponseDataType
	NUMBER   CreatePropertyResponseDataType
	STRING   CreatePropertyResponseDataType
	DATETIME CreatePropertyResponseDataType
	JSON     CreatePropertyResponseDataType
	BOOLEAN  CreatePropertyResponseDataType
	ARRAY    CreatePropertyResponseDataType
}

func GetCreatePropertyResponseDataTypeEnum() CreatePropertyResponseDataTypeEnum {
	return CreatePropertyResponseDataTypeEnum{
		INTEGER: CreatePropertyResponseDataType{
			value: "integer",
		},
		NUMBER: CreatePropertyResponseDataType{
			value: "number",
		},
		STRING: CreatePropertyResponseDataType{
			value: "string",
		},
		DATETIME: CreatePropertyResponseDataType{
			value: "datetime",
		},
		JSON: CreatePropertyResponseDataType{
			value: "json",
		},
		BOOLEAN: CreatePropertyResponseDataType{
			value: "boolean",
		},
		ARRAY: CreatePropertyResponseDataType{
			value: "array",
		},
	}
}

func (c CreatePropertyResponseDataType) Value() string {
	return c.value
}

func (c CreatePropertyResponseDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePropertyResponseDataType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreatePropertyResponseRequired struct {
	value int32
}

type CreatePropertyResponseRequiredEnum struct {
	E_0 CreatePropertyResponseRequired
	E_1 CreatePropertyResponseRequired
}

func GetCreatePropertyResponseRequiredEnum() CreatePropertyResponseRequiredEnum {
	return CreatePropertyResponseRequiredEnum{
		E_0: CreatePropertyResponseRequired{
			value: 0,
		}, E_1: CreatePropertyResponseRequired{
			value: 1,
		},
	}
}

func (c CreatePropertyResponseRequired) Value() int32 {
	return c.value
}

func (c CreatePropertyResponseRequired) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreatePropertyResponseRequired) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
