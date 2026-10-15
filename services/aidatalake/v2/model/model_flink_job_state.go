package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkJobState **参数解释**：作业状态的枚举字段。 **约束限制**：不涉及。 **取值范围**：   - CREATED：已创建。   - STARTING：启动中。   - RUNNING：运行中。   - CANCELLING：停止中。   - FAILED：已失败。   - CANCELLED：已停止。   - FINISHED：已完成。 **默认取值**：不涉及。
type FlinkJobState struct {
	value string
}

type FlinkJobStateEnum struct {
	CREATED    FlinkJobState
	STARTING   FlinkJobState
	RUNNING    FlinkJobState
	CANCELLING FlinkJobState
	FAILED     FlinkJobState
	CANCELLED  FlinkJobState
	FINISHED   FlinkJobState
}

func GetFlinkJobStateEnum() FlinkJobStateEnum {
	return FlinkJobStateEnum{
		CREATED: FlinkJobState{
			value: "CREATED",
		},
		STARTING: FlinkJobState{
			value: "STARTING",
		},
		RUNNING: FlinkJobState{
			value: "RUNNING",
		},
		CANCELLING: FlinkJobState{
			value: "CANCELLING",
		},
		FAILED: FlinkJobState{
			value: "FAILED",
		},
		CANCELLED: FlinkJobState{
			value: "CANCELLED",
		},
		FINISHED: FlinkJobState{
			value: "FINISHED",
		},
	}
}

func (c FlinkJobState) Value() string {
	return c.value
}

func (c FlinkJobState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkJobState) UnmarshalJSON(b []byte) error {
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
