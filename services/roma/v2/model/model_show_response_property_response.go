package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowResponsePropertyResponse struct {

	// 属性ID
	PropertyId *int32 `json:"property_id,omitempty" xml:"property_id"`

	// 属性名称，首位必须为字母，支持大小写字母，数字，中划线及下划线，长度2-50
	PropertyName *string `json:"property_name,omitempty" xml:"property_name"`

	// 属性描述，长度0-200
	Description *string `json:"description,omitempty" xml:"description"`

	// 属性数据类型，枚举值大小写敏感；number格式为数字，范围±1.0 x 10^-28 to ±7.9228 x 10^28；sting为字符串；integer为整数；datetime为时间，格式为yyyyMMddTHHmmss；json为自定义json格式
	DataType *ShowResponsePropertyResponseDataType `json:"data_type,omitempty" xml:"data_type"`

	// 是否必填 0-非必填 1-必填
	Required *ShowResponsePropertyResponseRequired `json:"required,omitempty" xml:"required"`

	// 最小值，当data_type为integer或number时有效
	Min *string `json:"min,omitempty" xml:"min"`

	// 最大值，当data_type为integer或number时有效
	Max *string `json:"max,omitempty" xml:"max"`

	// 步长，当data_type为integer或number时有效
	Step *string `json:"step,omitempty" xml:"step"`

	// 字符串最大长度，当data_type为string, datetime, json时有效
	MaxLength *int32 `json:"max_length,omitempty" xml:"max_length"`

	// 属性单位
	Unit *string `json:"unit,omitempty" xml:"unit"`

	// string的枚举值数组，使用逗号分隔
	EnumList       *string `json:"enum_list,omitempty" xml:"enum_list"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowResponsePropertyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowResponsePropertyResponse struct{}"
	}

	return strings.Join([]string{"ShowResponsePropertyResponse", string(data)}, " ")
}

type ShowResponsePropertyResponseDataType struct {
	value string
}

type ShowResponsePropertyResponseDataTypeEnum struct {
	INTEGER  ShowResponsePropertyResponseDataType
	NUMBER   ShowResponsePropertyResponseDataType
	STRING   ShowResponsePropertyResponseDataType
	DATETIME ShowResponsePropertyResponseDataType
	JSON     ShowResponsePropertyResponseDataType
}

func GetShowResponsePropertyResponseDataTypeEnum() ShowResponsePropertyResponseDataTypeEnum {
	return ShowResponsePropertyResponseDataTypeEnum{
		INTEGER: ShowResponsePropertyResponseDataType{
			value: "integer",
		},
		NUMBER: ShowResponsePropertyResponseDataType{
			value: "number",
		},
		STRING: ShowResponsePropertyResponseDataType{
			value: "string",
		},
		DATETIME: ShowResponsePropertyResponseDataType{
			value: "datetime",
		},
		JSON: ShowResponsePropertyResponseDataType{
			value: "json",
		},
	}
}

func (c ShowResponsePropertyResponseDataType) Value() string {
	return c.value
}

func (c ShowResponsePropertyResponseDataType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResponsePropertyResponseDataType) UnmarshalJSON(b []byte) error {
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

type ShowResponsePropertyResponseRequired struct {
	value int32
}

type ShowResponsePropertyResponseRequiredEnum struct {
	E_0 ShowResponsePropertyResponseRequired
	E_1 ShowResponsePropertyResponseRequired
}

func GetShowResponsePropertyResponseRequiredEnum() ShowResponsePropertyResponseRequiredEnum {
	return ShowResponsePropertyResponseRequiredEnum{
		E_0: ShowResponsePropertyResponseRequired{
			value: 0,
		}, E_1: ShowResponsePropertyResponseRequired{
			value: 1,
		},
	}
}

func (c ShowResponsePropertyResponseRequired) Value() int32 {
	return c.value
}

func (c ShowResponsePropertyResponseRequired) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowResponsePropertyResponseRequired) UnmarshalJSON(b []byte) error {
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
