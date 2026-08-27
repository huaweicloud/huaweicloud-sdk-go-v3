package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ApiType API协议类型（OpenAI/Anthropic）。
type ApiType struct {
	value string
}

type ApiTypeEnum struct {
	OPEN_AI   ApiType
	ANTHROPIC ApiType
}

func GetApiTypeEnum() ApiTypeEnum {
	return ApiTypeEnum{
		OPEN_AI: ApiType{
			value: "OpenAI",
		},
		ANTHROPIC: ApiType{
			value: "Anthropic",
		},
	}
}

func (c ApiType) Value() string {
	return c.value
}

func (c ApiType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ApiType) UnmarshalJSON(b []byte) error {
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
