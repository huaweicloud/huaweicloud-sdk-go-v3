package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ReqParam struct {

	// 参数名称。 由英文字母、数字、中划线、下划线、英文句号组成，且只能以英文开头。
	Name string `json:"name" xml:"name"`

	// 参数类型
	Type ReqParamType `json:"type" xml:"type"`

	// 参数位置
	Location ReqParamLocation `json:"location" xml:"location"`

	// 参数默认值
	DefaultValue *string `json:"default_value,omitempty" xml:"default_value"`

	// 参数示例值
	SampleValue *string `json:"sample_value,omitempty" xml:"sample_value"`

	// 是否必须 - 1：是 - 2：否  location为PATH时，required默认为1，其他场景required默认为2
	Required *ReqParamRequired `json:"required,omitempty" xml:"required"`

	// 是否开启校验 - 1：开启校验 - 2：不开启校验
	ValidEnable *ReqParamValidEnable `json:"valid_enable,omitempty" xml:"valid_enable"`

	// 描述信息。 > 中文字符必须为UTF-8或者unicode编码。
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
	PassThrough *ReqParamPassThrough `json:"pass_through,omitempty" xml:"pass_through"`

	// 参数编号
	Id *string `json:"id,omitempty" xml:"id"`
}

func (o ReqParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReqParam struct{}"
	}

	return strings.Join([]string{"ReqParam", string(data)}, " ")
}

type ReqParamType struct {
	value string
}

type ReqParamTypeEnum struct {
	STRING ReqParamType
	NUMBER ReqParamType
}

func GetReqParamTypeEnum() ReqParamTypeEnum {
	return ReqParamTypeEnum{
		STRING: ReqParamType{
			value: "STRING",
		},
		NUMBER: ReqParamType{
			value: "NUMBER",
		},
	}
}

func (c ReqParamType) Value() string {
	return c.value
}

func (c ReqParamType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqParamType) UnmarshalJSON(b []byte) error {
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

type ReqParamLocation struct {
	value string
}

type ReqParamLocationEnum struct {
	PATH   ReqParamLocation
	QUERY  ReqParamLocation
	HEADER ReqParamLocation
}

func GetReqParamLocationEnum() ReqParamLocationEnum {
	return ReqParamLocationEnum{
		PATH: ReqParamLocation{
			value: "PATH",
		},
		QUERY: ReqParamLocation{
			value: "QUERY",
		},
		HEADER: ReqParamLocation{
			value: "HEADER",
		},
	}
}

func (c ReqParamLocation) Value() string {
	return c.value
}

func (c ReqParamLocation) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqParamLocation) UnmarshalJSON(b []byte) error {
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

type ReqParamRequired struct {
	value int32
}

type ReqParamRequiredEnum struct {
	E_1 ReqParamRequired
	E_2 ReqParamRequired
}

func GetReqParamRequiredEnum() ReqParamRequiredEnum {
	return ReqParamRequiredEnum{
		E_1: ReqParamRequired{
			value: 1,
		}, E_2: ReqParamRequired{
			value: 2,
		},
	}
}

func (c ReqParamRequired) Value() int32 {
	return c.value
}

func (c ReqParamRequired) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqParamRequired) UnmarshalJSON(b []byte) error {
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

type ReqParamValidEnable struct {
	value int32
}

type ReqParamValidEnableEnum struct {
	E_1 ReqParamValidEnable
	E_2 ReqParamValidEnable
}

func GetReqParamValidEnableEnum() ReqParamValidEnableEnum {
	return ReqParamValidEnableEnum{
		E_1: ReqParamValidEnable{
			value: 1,
		}, E_2: ReqParamValidEnable{
			value: 2,
		},
	}
}

func (c ReqParamValidEnable) Value() int32 {
	return c.value
}

func (c ReqParamValidEnable) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqParamValidEnable) UnmarshalJSON(b []byte) error {
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

type ReqParamPassThrough struct {
	value int32
}

type ReqParamPassThroughEnum struct {
	E_1 ReqParamPassThrough
	E_2 ReqParamPassThrough
}

func GetReqParamPassThroughEnum() ReqParamPassThroughEnum {
	return ReqParamPassThroughEnum{
		E_1: ReqParamPassThrough{
			value: 1,
		}, E_2: ReqParamPassThrough{
			value: 2,
		},
	}
}

func (c ReqParamPassThrough) Value() int32 {
	return c.value
}

func (c ReqParamPassThrough) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ReqParamPassThrough) UnmarshalJSON(b []byte) error {
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
