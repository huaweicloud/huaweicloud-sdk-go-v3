package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LlmSourceEnum 使用的大语言模型来源 * LLM_CONFIG：大语言模型配置 * DEEP_SEEK：内置的DEEPSEEK
type LlmSourceEnum struct {
	value string
}

type LlmSourceEnumEnum struct {
	LLM_CONFIG LlmSourceEnum
	DEEP_SEEK  LlmSourceEnum
}

func GetLlmSourceEnumEnum() LlmSourceEnumEnum {
	return LlmSourceEnumEnum{
		LLM_CONFIG: LlmSourceEnum{
			value: "LLM_CONFIG",
		},
		DEEP_SEEK: LlmSourceEnum{
			value: "DEEP_SEEK",
		},
	}
}

func (c LlmSourceEnum) Value() string {
	return c.value
}

func (c LlmSourceEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LlmSourceEnum) UnmarshalJSON(b []byte) error {
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
