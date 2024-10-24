package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ModelType 模型的类型，可选值： - LLM_MODEL：大语言模型
type ModelType struct {
	value string
}

type ModelTypeEnum struct {
	LLM_MODEL ModelType
}

func GetModelTypeEnum() ModelTypeEnum {
	return ModelTypeEnum{
		LLM_MODEL: ModelType{
			value: "LLM_MODEL",
		},
	}
}

func (c ModelType) Value() string {
	return c.value
}

func (c ModelType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ModelType) UnmarshalJSON(b []byte) error {
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
