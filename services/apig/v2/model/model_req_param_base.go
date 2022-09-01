package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ReqParamBase struct {

	// 参数名称。 长度为1 ~ 32位的字符串，字符串由英文字母、数字、中划线、下划线、英文句号组成，且只能以英文开头。
	Name string `json:"name" xml:"name"`

	// 参数类型
	Type ReqParamBaseType `json:"type" xml:"type"`

	// 参数位置
	Location ReqParamBaseLocation `json:"location" xml:"location"`

	// 参数默认值
	DefaultValue *string `json:"default_value,omitempty" xml:"default_value"`

	// 参数示例值
	SampleValue *string `json:"sample_value,omitempty" xml:"sample_value"`

	// 是否必须 - 1：是 - 2：否  location为PATH时，required默认为1，其他场景required默认为2
	Required *ReqParamBaseRequired `json:"required,omitempty" xml:"required"`

	// 是否开启校验 - 1：开启校验 - 2：不开启校验
	ValidEnable *ReqParamBaseValidEnable `json:"valid_enable,omitempty" xml:"valid_enable"`

	// 描述信息。长度不超过255个字符 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty" xml:"remark"`

	// 参数枚举值
	Enumerations *string `json:"enumerations,omitempty" xml:"enumerations"`

	// 参数最小值  参数类型为NUMBER时有效
	MinNum *int32 `json:"min_num,omitempty" xml:"min_num"`

	// 参数最大值  参数类型为NUMBER时有效
	MaxNum *int32 `json:"max_num,omitempty" xml:"max_num"`

	// 参数最小长度  参数类型为STRING时有效
	MinSize *int32 `json:"min_size,omitempty" xml:"min_size"`

	// 参数最大长度  参数类型为STRING时有效
	MaxSize *int32 `json:"max_size,omitempty" xml:"max_size"`

	// 正则校验规则  暂不支持
	Regular *string `json:"regular,omitempty" xml:"regular"`

	// JSON校验规则  暂不支持
	JsonSchema *string `json:"json_schema,omitempty" xml:"json_schema"`

	// 是否透传 - 1：是 - 2：否
	PassThrough *ReqParamBasePassThrough `json:"pass_through,omitempty" xml:"pass_through"`
}

func (o ReqParamBase) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqParamBase struct{}"
	}

	return strings.Join([]string{"ReqParamBase", string(data)}, " ")
}

type ReqParamBaseType struct {
	value string
}

type ReqParamBaseTypeEnum struct {
	STRING ReqParamBaseType
	NUMBER ReqParamBaseType
}

func GetReqParamBaseTypeEnum() ReqParamBaseTypeEnum {
	return ReqParamBaseTypeEnum{
		STRING: ReqParamBaseType{
			value: "STRING",
		},
		NUMBER: ReqParamBaseType{
			value: "NUMBER",
		},
	}
}

func (c ReqParamBaseType) Value() string {
	return c.value
}

func (c ReqParamBaseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqParamBaseType) UnmarshalJSON(b []byte) error {
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

type ReqParamBaseLocation struct {
	value string
}

type ReqParamBaseLocationEnum struct {
	PATH   ReqParamBaseLocation
	QUERY  ReqParamBaseLocation
	HEADER ReqParamBaseLocation
}

func GetReqParamBaseLocationEnum() ReqParamBaseLocationEnum {
	return ReqParamBaseLocationEnum{
		PATH: ReqParamBaseLocation{
			value: "PATH",
		},
		QUERY: ReqParamBaseLocation{
			value: "QUERY",
		},
		HEADER: ReqParamBaseLocation{
			value: "HEADER",
		},
	}
}

func (c ReqParamBaseLocation) Value() string {
	return c.value
}

func (c ReqParamBaseLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqParamBaseLocation) UnmarshalJSON(b []byte) error {
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

type ReqParamBaseRequired struct {
	value int32
}

type ReqParamBaseRequiredEnum struct {
	E_1 ReqParamBaseRequired
	E_2 ReqParamBaseRequired
}

func GetReqParamBaseRequiredEnum() ReqParamBaseRequiredEnum {
	return ReqParamBaseRequiredEnum{
		E_1: ReqParamBaseRequired{
			value: 1,
		}, E_2: ReqParamBaseRequired{
			value: 2,
		},
	}
}

func (c ReqParamBaseRequired) Value() int32 {
	return c.value
}

func (c ReqParamBaseRequired) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqParamBaseRequired) UnmarshalJSON(b []byte) error {
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

type ReqParamBaseValidEnable struct {
	value int32
}

type ReqParamBaseValidEnableEnum struct {
	E_1 ReqParamBaseValidEnable
	E_2 ReqParamBaseValidEnable
}

func GetReqParamBaseValidEnableEnum() ReqParamBaseValidEnableEnum {
	return ReqParamBaseValidEnableEnum{
		E_1: ReqParamBaseValidEnable{
			value: 1,
		}, E_2: ReqParamBaseValidEnable{
			value: 2,
		},
	}
}

func (c ReqParamBaseValidEnable) Value() int32 {
	return c.value
}

func (c ReqParamBaseValidEnable) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqParamBaseValidEnable) UnmarshalJSON(b []byte) error {
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

type ReqParamBasePassThrough struct {
	value int32
}

type ReqParamBasePassThroughEnum struct {
	E_1 ReqParamBasePassThrough
	E_2 ReqParamBasePassThrough
}

func GetReqParamBasePassThroughEnum() ReqParamBasePassThroughEnum {
	return ReqParamBasePassThroughEnum{
		E_1: ReqParamBasePassThrough{
			value: 1,
		}, E_2: ReqParamBasePassThrough{
			value: 2,
		},
	}
}

func (c ReqParamBasePassThrough) Value() int32 {
	return c.value
}

func (c ReqParamBasePassThrough) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqParamBasePassThrough) UnmarshalJSON(b []byte) error {
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
