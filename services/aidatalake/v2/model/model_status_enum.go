package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// StatusEnum **参数解释**：运行状态。 **约束限制**：不涉及。 **取值范围**：   - QUEUED：排队中。   - PENDING：待处理。   - RUNNING：运行中。   - CANCELING：取消中。   - CANCELED：已取消。   - FAILED：失败。   - QUEUED_TIMEOUT：排队超时。   - RUNNING_TIMEOUT：运行超时。   - SUCCEEDED：成功 **默认取值**：不涉及。
type StatusEnum struct {
	value string
}

type StatusEnumEnum struct {
	QUEUED          StatusEnum
	PENDING         StatusEnum
	RUNNING         StatusEnum
	CANCELING       StatusEnum
	CANCELED        StatusEnum
	FAILED          StatusEnum
	QUEUED_TIMEOUT  StatusEnum
	RUNNING_TIMEOUT StatusEnum
	SUCCEEDED       StatusEnum
}

func GetStatusEnumEnum() StatusEnumEnum {
	return StatusEnumEnum{
		QUEUED: StatusEnum{
			value: "QUEUED",
		},
		PENDING: StatusEnum{
			value: "PENDING",
		},
		RUNNING: StatusEnum{
			value: "RUNNING",
		},
		CANCELING: StatusEnum{
			value: "CANCELING",
		},
		CANCELED: StatusEnum{
			value: "CANCELED",
		},
		FAILED: StatusEnum{
			value: "FAILED",
		},
		QUEUED_TIMEOUT: StatusEnum{
			value: "QUEUED_TIMEOUT",
		},
		RUNNING_TIMEOUT: StatusEnum{
			value: "RUNNING_TIMEOUT",
		},
		SUCCEEDED: StatusEnum{
			value: "SUCCEEDED",
		},
	}
}

func (c StatusEnum) Value() string {
	return c.value
}

func (c StatusEnum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *StatusEnum) UnmarshalJSON(b []byte) error {
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
