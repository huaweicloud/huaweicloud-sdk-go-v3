package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Status **参数解释**: 状态 - ENABLED 启用 - DISABLED 禁用  **约束限制** 不涉及 **取值范围**: - ENABLED - DISABLED  **默认值** 不涉及
type Status struct {
	value string
}

type StatusEnum struct {
	ENABLED  Status
	DISABLED Status
}

func GetStatusEnum() StatusEnum {
	return StatusEnum{
		ENABLED: Status{
			value: "ENABLED",
		},
		DISABLED: Status{
			value: "DISABLED",
		},
	}
}

func (c Status) Value() string {
	return c.value
}

func (c Status) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *Status) UnmarshalJSON(b []byte) error {
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
