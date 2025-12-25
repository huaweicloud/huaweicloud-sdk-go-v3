package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ProcessError **参数解释**: 处理错误 - NONE 无  **约束限制** 不涉及 **取值范围**: - NONE  **默认值** 不涉及
type ProcessError struct {
	value string
}

type ProcessErrorEnum struct {
	NONE ProcessError
}

func GetProcessErrorEnum() ProcessErrorEnum {
	return ProcessErrorEnum{
		NONE: ProcessError{
			value: "NONE",
		},
	}
}

func (c ProcessError) Value() string {
	return c.value
}

func (c ProcessError) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ProcessError) UnmarshalJSON(b []byte) error {
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
