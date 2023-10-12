package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type EnvRespParams struct {

	// 变量名称
	Name *string `json:"name,omitempty"`

	// 变量类型
	Type *EnvRespParamsType `json:"type,omitempty"`

	// 变量值
	Value *string `json:"value,omitempty"`

	// 描述信息
	Desc *string `json:"desc,omitempty"`
}

func (o EnvRespParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EnvRespParams struct{}"
	}

	return strings.Join([]string{"EnvRespParams", string(data)}, " ")
}

type EnvRespParamsType struct {
	value string
}

type EnvRespParamsTypeEnum struct {
	VARIABLE EnvRespParamsType
	CONSTANT EnvRespParamsType
}

func GetEnvRespParamsTypeEnum() EnvRespParamsTypeEnum {
	return EnvRespParamsTypeEnum{
		VARIABLE: EnvRespParamsType{
			value: "variable",
		},
		CONSTANT: EnvRespParamsType{
			value: "constant",
		},
	}
}

func (c EnvRespParamsType) Value() string {
	return c.value
}

func (c EnvRespParamsType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *EnvRespParamsType) UnmarshalJSON(b []byte) error {
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
