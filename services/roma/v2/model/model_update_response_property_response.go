package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateResponsePropertyResponse Response Object
type UpdateResponsePropertyResponse struct {

	// 属性ID
	PropertyId *int32 `json:"property_id,omitempty"`

	// 属性名称，首位必须为字母，支持大小写字母，数字，中划线及下划线，长度2-50
	PropertyName *string `json:"property_name,omitempty"`

	// 属性描述，长度0-200
	Description *string `json:"description,omitempty"`

	// 属性数据类型，boolean枚举值大小写敏感；number格式为数字，范围±1.0 x 10^-28 to ±7.9228 x 10^28；string为字符串；integer为整数；datetime为时间，格式为yyyyMMddTHHmmss；json为自定义json格式; array为数组类型
	DataType *UpdateResponsePropertyResponseDataType `json:"data_type,omitempty"`

	// 是否必填 0-非必填 1-必填
	Required *UpdateResponsePropertyResponseRequired `json:"required,omitempty"`

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

func (o UpdateResponsePropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateResponsePropertyResponse struct{}"
	}

	return strings.Join([]string{"UpdateResponsePropertyResponse", string(data)}, " ")
}

type UpdateResponsePropertyResponseDataType struct {
	value string
}

type UpdateResponsePropertyResponseDataTypeEnum struct {
	INTEGER  UpdateResponsePropertyResponseDataType
	NUMBER   UpdateResponsePropertyResponseDataType
	STRING   UpdateResponsePropertyResponseDataType
	DATETIME UpdateResponsePropertyResponseDataType
	JSON     UpdateResponsePropertyResponseDataType
	BOOLEAN  UpdateResponsePropertyResponseDataType
	ARRAY    UpdateResponsePropertyResponseDataType
}

func GetUpdateResponsePropertyResponseDataTypeEnum() UpdateResponsePropertyResponseDataTypeEnum {
	return UpdateResponsePropertyResponseDataTypeEnum{
		INTEGER: UpdateResponsePropertyResponseDataType{
			value: "integer",
		},
		NUMBER: UpdateResponsePropertyResponseDataType{
			value: "number",
		},
		STRING: UpdateResponsePropertyResponseDataType{
			value: "string",
		},
		DATETIME: UpdateResponsePropertyResponseDataType{
			value: "datetime",
		},
		JSON: UpdateResponsePropertyResponseDataType{
			value: "json",
		},
		BOOLEAN: UpdateResponsePropertyResponseDataType{
			value: "boolean",
		},
		ARRAY: UpdateResponsePropertyResponseDataType{
			value: "array",
		},
	}
}

func (c UpdateResponsePropertyResponseDataType) Value() string {
	return c.value
}

func (c UpdateResponsePropertyResponseDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateResponsePropertyResponseDataType) UnmarshalJSON(b []byte) error {
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

type UpdateResponsePropertyResponseRequired struct {
	value int32
}

type UpdateResponsePropertyResponseRequiredEnum struct {
	E_0 UpdateResponsePropertyResponseRequired
	E_1 UpdateResponsePropertyResponseRequired
}

func GetUpdateResponsePropertyResponseRequiredEnum() UpdateResponsePropertyResponseRequiredEnum {
	return UpdateResponsePropertyResponseRequiredEnum{
		E_0: UpdateResponsePropertyResponseRequired{
			value: 0,
		}, E_1: UpdateResponsePropertyResponseRequired{
			value: 1,
		},
	}
}

func (c UpdateResponsePropertyResponseRequired) Value() int32 {
	return c.value
}

func (c UpdateResponsePropertyResponseRequired) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateResponsePropertyResponseRequired) UnmarshalJSON(b []byte) error {
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
