package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Mode **参数解释**: 模式 - COUNT 计数  **约束限制** 不涉及 **取值范围**: - COUNT  **默认值** 不涉及
type Mode struct {
	value string
}

type ModeEnum struct {
	COUNT Mode
}

func GetModeEnum() ModeEnum {
	return ModeEnum{
		COUNT: Mode{
			value: "COUNT",
		},
	}
}

func (c Mode) Value() string {
	return c.value
}

func (c Mode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Mode) UnmarshalJSON(b []byte) error {
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
