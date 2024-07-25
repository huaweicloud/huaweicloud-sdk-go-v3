package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ReqParam struct {

	// 参数名称。 长度为1 ~ 32位的字符串，字符串由英文字母、数字、中划线、下划线、英文句号组成，且只能以英文开头。
	Name string `json:"name"`

	// 参数类型
	Type ReqParamType `json:"type"`

	// 参数位置
	Location ReqParamLocation `json:"location"`

	// 参数默认值
	DefaultValue *string `json:"default_value,omitempty"`

	// 参数示例值
	SampleValue *string `json:"sample_value,omitempty"`

	// 是否必须 - 1：是 - 2：否  location为PATH时，required默认为1，其他场景required默认为2
	Required *ReqParamRequired `json:"required,omitempty"`

	// 是否开启校验 - 1：开启校验 - 2：不开启校验
	ValidEnable *ReqParamValidEnable `json:"valid_enable,omitempty"`

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
	PassThrough *ReqParamPassThrough `json:"pass_through,omitempty"`

	// 请求参数匹配编排规则的生效优先级与列表顺序保持一致，列表中靠前的配置匹配优先级较高； 如果编配规则列表中包含none_value类型的规则，则none_value类型的规则优先级最高，至多绑定一个none_value类型的规则； 如果编排规则列表中包含default类型的规则，则default类型的规则优先级最低，至多绑定一个default类型的规则； 当编排规则为预处理策略时，该规则不能作为除default以外的最后一个编排规则； 每个API仅允许选择一个参数绑定编排规则，且编排规则不能重复，支持绑定的编排规则数量有配额限制，具体请参见产品介绍的“配额说明”章节。
	Orchestrations *[]string `json:"orchestrations,omitempty"`

	// 参数编号
	Id *string `json:"id,omitempty"`
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

type ReqParamLocation struct {
	value string
}

type ReqParamLocationEnum struct {
	PATH   ReqParamLocation
	QUERY  ReqParamLocation
	HEADER ReqParamLocation
	COOKIE ReqParamLocation
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
		COOKIE: ReqParamLocation{
			value: "COOKIE",
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
