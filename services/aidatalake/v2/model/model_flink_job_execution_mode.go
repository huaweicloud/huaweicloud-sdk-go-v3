package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkJobExecutionMode **参数解释**：执行模式枚举。 **约束限制**：不涉及。 **取值范围**：   - STREAMING：流模式。   - BATCH：批模式。 **默认取值**：不涉及。
type FlinkJobExecutionMode struct {
	value string
}

type FlinkJobExecutionModeEnum struct {
	STREAMING FlinkJobExecutionMode
	BATCH     FlinkJobExecutionMode
}

func GetFlinkJobExecutionModeEnum() FlinkJobExecutionModeEnum {
	return FlinkJobExecutionModeEnum{
		STREAMING: FlinkJobExecutionMode{
			value: "STREAMING",
		},
		BATCH: FlinkJobExecutionMode{
			value: "BATCH",
		},
	}
}

func (c FlinkJobExecutionMode) Value() string {
	return c.value
}

func (c FlinkJobExecutionMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkJobExecutionMode) UnmarshalJSON(b []byte) error {
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
