package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkJobCancelStrategy **参数解释**：作业停止策略。 **约束限制**：不涉及。 **取值范围**：   - NONE：立即停止，不保存状态。   - STOP_WITH_SAVEPOINT：停止前创建快照。   - STOP_WITH_DRAIN：Drain模式停止，先触发所有窗口计算并发送水位线，再创建快照停止。   - FORCE_CANCEL：强制取消，仅当作业已处于CANCELLING状态时可再次触发取消。 **默认取值**：不涉及。
type FlinkJobCancelStrategy struct {
	value string
}

type FlinkJobCancelStrategyEnum struct {
	NONE                FlinkJobCancelStrategy
	STOP_WITH_SAVEPOINT FlinkJobCancelStrategy
	STOP_WITH_DRAIN     FlinkJobCancelStrategy
	FORCE_CANCEL        FlinkJobCancelStrategy
}

func GetFlinkJobCancelStrategyEnum() FlinkJobCancelStrategyEnum {
	return FlinkJobCancelStrategyEnum{
		NONE: FlinkJobCancelStrategy{
			value: "NONE",
		},
		STOP_WITH_SAVEPOINT: FlinkJobCancelStrategy{
			value: "STOP_WITH_SAVEPOINT",
		},
		STOP_WITH_DRAIN: FlinkJobCancelStrategy{
			value: "STOP_WITH_DRAIN",
		},
		FORCE_CANCEL: FlinkJobCancelStrategy{
			value: "FORCE_CANCEL",
		},
	}
}

func (c FlinkJobCancelStrategy) Value() string {
	return c.value
}

func (c FlinkJobCancelStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkJobCancelStrategy) UnmarshalJSON(b []byte) error {
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
