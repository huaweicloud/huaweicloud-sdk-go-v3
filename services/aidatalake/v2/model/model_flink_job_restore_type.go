package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// FlinkJobRestoreType **参数解释**：Flink作业恢复类型枚举。 **约束限制**：不涉及。 **取值范围**：   - NONE：无状态启动。   - FROM_SAVEPOINT_ID：从指定快照恢复。   - FROM_CHECKPOINT_PATH：从指定Checkpoint路径恢复。   - FROM_SPECIFIED_TIMESTAMP：从指定时间戳恢复。 **默认取值**：NONE。
type FlinkJobRestoreType struct {
	value string
}

type FlinkJobRestoreTypeEnum struct {
	NONE                     FlinkJobRestoreType
	FROM_SAVEPOINT_ID        FlinkJobRestoreType
	FROM_CHECKPOINT_PATH     FlinkJobRestoreType
	FROM_SPECIFIED_TIMESTAMP FlinkJobRestoreType
}

func GetFlinkJobRestoreTypeEnum() FlinkJobRestoreTypeEnum {
	return FlinkJobRestoreTypeEnum{
		NONE: FlinkJobRestoreType{
			value: "NONE",
		},
		FROM_SAVEPOINT_ID: FlinkJobRestoreType{
			value: "FROM_SAVEPOINT_ID",
		},
		FROM_CHECKPOINT_PATH: FlinkJobRestoreType{
			value: "FROM_CHECKPOINT_PATH",
		},
		FROM_SPECIFIED_TIMESTAMP: FlinkJobRestoreType{
			value: "FROM_SPECIFIED_TIMESTAMP",
		},
	}
}

func (c FlinkJobRestoreType) Value() string {
	return c.value
}

func (c FlinkJobRestoreType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *FlinkJobRestoreType) UnmarshalJSON(b []byte) error {
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
