package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkSavepointState **参数解释**：Savepoint状态枚举。 **约束限制**：不涉及。 **取值范围**：   - TRIGGER_PENDING：排队等待执行。   - IN_PROGRESS：执行中。   - COMPLETED：触发成功。   - FAILED：触发失败。   - ABANDONED：放弃/失效。 **默认取值**：不涉及。
type FlinkSavepointState struct {
	value string
}

type FlinkSavepointStateEnum struct {
	TRIGGER_PENDING FlinkSavepointState
	IN_PROGRESS     FlinkSavepointState
	COMPLETED       FlinkSavepointState
	FAILED          FlinkSavepointState
	ABANDONED       FlinkSavepointState
}

func GetFlinkSavepointStateEnum() FlinkSavepointStateEnum {
	return FlinkSavepointStateEnum{
		TRIGGER_PENDING: FlinkSavepointState{
			value: "TRIGGER_PENDING",
		},
		IN_PROGRESS: FlinkSavepointState{
			value: "IN_PROGRESS",
		},
		COMPLETED: FlinkSavepointState{
			value: "COMPLETED",
		},
		FAILED: FlinkSavepointState{
			value: "FAILED",
		},
		ABANDONED: FlinkSavepointState{
			value: "ABANDONED",
		},
	}
}

func (c FlinkSavepointState) Value() string {
	return c.value
}

func (c FlinkSavepointState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkSavepointState) UnmarshalJSON(b []byte) error {
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
