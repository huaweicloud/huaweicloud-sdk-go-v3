package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type UpdateRequestPropertyResponse struct {

	// 属性ID
	PropertyId *int32 `json:"property_id,omitempty"`

	// 属性名称，首位必须为字母，支持大小写字母，数字，中划线及下划线，长度2-50
	PropertyName *string `json:"property_name,omitempty"`

	// 属性描述，长度0-200
	Description *string `json:"description,omitempty"`

	// 属性数据类型，boolean枚举值大小写敏感；number格式为数字，范围±1.0 x 10^-28 to ±7.9228 x 10^28；string为字符串；integer为整数；datetime为时间，格式为yyyyMMddTHHmmss；json为自定义json格式; array为数组类型
	DataType *UpdateRequestPropertyResponseDataType `json:"data_type,omitempty"`

	// 是否必填 0-非必填 1-必填
	Required *UpdateRequestPropertyResponseRequired `json:"required,omitempty"`

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
	EnumDict       *interface{} `json:"enum_dict,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o UpdateRequestPropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRequestPropertyResponse struct{}"
	}

	return strings.Join([]string{"UpdateRequestPropertyResponse", string(data)}, " ")
}

type UpdateRequestPropertyResponseDataType struct {
	value string
}

type UpdateRequestPropertyResponseDataTypeEnum struct {
	INTEGER  UpdateRequestPropertyResponseDataType
	NUMBER   UpdateRequestPropertyResponseDataType
	STRING   UpdateRequestPropertyResponseDataType
	DATETIME UpdateRequestPropertyResponseDataType
	JSON     UpdateRequestPropertyResponseDataType
	BOOLEAN  UpdateRequestPropertyResponseDataType
	ARRAY    UpdateRequestPropertyResponseDataType
}

func GetUpdateRequestPropertyResponseDataTypeEnum() UpdateRequestPropertyResponseDataTypeEnum {
	return UpdateRequestPropertyResponseDataTypeEnum{
		INTEGER: UpdateRequestPropertyResponseDataType{
			value: "integer",
		},
		NUMBER: UpdateRequestPropertyResponseDataType{
			value: "number",
		},
		STRING: UpdateRequestPropertyResponseDataType{
			value: "string",
		},
		DATETIME: UpdateRequestPropertyResponseDataType{
			value: "datetime",
		},
		JSON: UpdateRequestPropertyResponseDataType{
			value: "json",
		},
		BOOLEAN: UpdateRequestPropertyResponseDataType{
			value: "boolean",
		},
		ARRAY: UpdateRequestPropertyResponseDataType{
			value: "array",
		},
	}
}

func (c UpdateRequestPropertyResponseDataType) Value() string {
	return c.value
}

func (c UpdateRequestPropertyResponseDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRequestPropertyResponseDataType) UnmarshalJSON(b []byte) error {
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

type UpdateRequestPropertyResponseRequired struct {
	value int32
}

type UpdateRequestPropertyResponseRequiredEnum struct {
	E_0 UpdateRequestPropertyResponseRequired
	E_1 UpdateRequestPropertyResponseRequired
}

func GetUpdateRequestPropertyResponseRequiredEnum() UpdateRequestPropertyResponseRequiredEnum {
	return UpdateRequestPropertyResponseRequiredEnum{
		E_0: UpdateRequestPropertyResponseRequired{
			value: 0,
		}, E_1: UpdateRequestPropertyResponseRequired{
			value: 1,
		},
	}
}

func (c UpdateRequestPropertyResponseRequired) Value() int32 {
	return c.value
}

func (c UpdateRequestPropertyResponseRequired) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateRequestPropertyResponseRequired) UnmarshalJSON(b []byte) error {
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
