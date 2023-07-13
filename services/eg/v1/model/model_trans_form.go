package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TransForm 订阅的事件目标转换规则
type TransForm struct {

	// 转换规则类型
	Type TransFormType `json:"type"`

	// 常量类型规则时，字段为常量内容定义； 变量类型规则时，为变量定义，内容必须为JsonObject字符串。 变量最多支持100个，且不支持嵌套结构定义； 变量名由字母、数字、点、下划线和中划线组成，必须字母或数字开头不能以HC.开头，长度不超过64个字符； 变量值表达式支持常量或JsonPath表达式，字符串长度不超过1024个字符。
	Value *string `json:"value,omitempty"`

	// 变量类型规则时，规则内容的模板定义，支持对已定义变量的引用。
	Template *string `json:"template,omitempty"`
}

func (o TransForm) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TransForm struct{}"
	}

	return strings.Join([]string{"TransForm", string(data)}, " ")
}

type TransFormType struct {
	value string
}

type TransFormTypeEnum struct {
	ORIGINAL TransFormType
	CONSTANT TransFormType
	VARIABLE TransFormType
}

func GetTransFormTypeEnum() TransFormTypeEnum {
	return TransFormTypeEnum{
		ORIGINAL: TransFormType{
			value: "ORIGINAL",
		},
		CONSTANT: TransFormType{
			value: "CONSTANT",
		},
		VARIABLE: TransFormType{
			value: "VARIABLE",
		},
	}
}

func (c TransFormType) Value() string {
	return c.value
}

func (c TransFormType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TransFormType) UnmarshalJSON(b []byte) error {
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
