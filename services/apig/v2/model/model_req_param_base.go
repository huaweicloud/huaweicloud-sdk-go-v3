package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ReqParamBase struct {

	// 参数名称。 长度为1 ~ 32位的字符串，字符串由英文字母、数字、中划线、下划线、英文句号组成，且只能以英文开头。
	Name string `json:"name"`

	// 参数类型
	Type ReqParamBaseType `json:"type"`

	// 参数位置
	Location ReqParamBaseLocation `json:"location"`

	// 参数默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 参数示例值
	SampleValue *string `json:"sample_value,omitempty"`

	// 是否必须 - 1：是 - 2：否  location为PATH时，required默认为1，其他场景required默认为2
	Required *ReqParamBaseRequired `json:"required,omitempty"`

	// 是否开启校验 - 1：开启校验 - 2：不开启校验
	ValidEnable *ReqParamBaseValidEnable `json:"valid_enable,omitempty"`

	// 描述信息。长度不超过255个字符 > 中文字符必须为UTF-8或者unicode编码。
	Remark *string `json:"remark,omitempty"`

	// 参数枚举值
	Enumerations *string `json:"enumerations,omitempty"`

	// 参数最小值  参数类型为NUMBER时有效
	MinNum *int32 `json:"min_num,omitempty"`

	// 参数最大值  参数类型为NUMBER时有效
	MaxNum *int32 `json:"max_num,omitempty"`

	// 参数最小长度  参数类型为STRING时有效
	MinSize *int32 `json:"min_size,omitempty"`

	// 参数最大长度  参数类型为STRING时有效
	MaxSize *int32 `json:"max_size,omitempty"`

	// 正则校验规则  暂不支持
	Regular *string `json:"regular,omitempty"`

	// JSON校验规则  暂不支持
	JsonSchema *string `json:"json_schema,omitempty"`

	// 是否透传 - 1：是 - 2：否
	PassThrough *ReqParamBasePassThrough `json:"pass_through,omitempty"`

	// 请求参数匹配编排规则的生效优先级与列表顺序保持一致，列表中靠前的配置匹配优先级较高； 如果编配规则列表中包含none_value类型的规则，则none_value类型的规则优先级最高，至多绑定一个none_value类型的规则； 如果编排规则列表中包含default类型的规则，则default类型的规则优先级最低，至多绑定一个default类型的规则； 当编排规则为预处理策略时，该规则不能作为除default以外的最后一个编排规则； 每个API仅允许选择一个参数绑定编排规则，且编排规则不能重复，支持绑定的编排规则数量有配额限制，具体请参见产品介绍的“配额说明”章节。
	Orchestrations *[]string `json:"orchestrations,omitempty"`
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

type ReqParamBaseLocation struct {
	value string
}

type ReqParamBaseLocationEnum struct {
	PATH   ReqParamBaseLocation
	QUERY  ReqParamBaseLocation
	HEADER ReqParamBaseLocation
	COOKIE ReqParamBaseLocation
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
		COOKIE: ReqParamBaseLocation{
			value: "COOKIE",
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
