package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type CheckPoint struct {

	// 运算符
	Comparison *string `json:"comparison,omitempty"`

	// 描述信息
	Description *string `json:"description,omitempty"`

	// 响应提取时要使用什么方法处理参数
	FunctionArg *string `json:"function_arg,omitempty"`

	// 响应提取时要使用什么方法处理参数
	FunctionType *CheckPointFunctionType `json:"function_type,omitempty"`

	// 属性名称
	Property *string `json:"property,omitempty"`

	// 值
	Value *string `json:"value,omitempty"`
}

func (o CheckPoint) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CheckPoint struct{}"
	}

	return strings.Join([]string{"CheckPoint", string(data)}, " ")
}

type CheckPointFunctionType struct {
	value string
}

type CheckPointFunctionTypeEnum struct {
	REGEX     CheckPointFunctionType
	SUBSTRING CheckPointFunctionType
}

func GetCheckPointFunctionTypeEnum() CheckPointFunctionTypeEnum {
	return CheckPointFunctionTypeEnum{
		REGEX: CheckPointFunctionType{
			value: "REGEX",
		},
		SUBSTRING: CheckPointFunctionType{
			value: "SUBSTRING",
		},
	}
}

func (c CheckPointFunctionType) Value() string {
	return c.value
}

func (c CheckPointFunctionType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CheckPointFunctionType) UnmarshalJSON(b []byte) error {
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
