package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ExecuteInstancesBatchInfo struct {

	// 批次索引，从1开始，未校验：最大值
	BatchIndex int32 `json:"batch_index"`

	// 目标节点列表
	TargetInstances []ExecuteResourceInstance `json:"target_instances"`

	// 暂停继续策略
	RotationStrategy ExecuteInstancesBatchInfoRotationStrategy `json:"rotation_strategy"`
}

func (o ExecuteInstancesBatchInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExecuteInstancesBatchInfo struct{}"
	}

	return strings.Join([]string{"ExecuteInstancesBatchInfo", string(data)}, " ")
}

type ExecuteInstancesBatchInfoRotationStrategy struct {
	value string
}

type ExecuteInstancesBatchInfoRotationStrategyEnum struct {
	CONTINUE ExecuteInstancesBatchInfoRotationStrategy
	PAUSE    ExecuteInstancesBatchInfoRotationStrategy
}

func GetExecuteInstancesBatchInfoRotationStrategyEnum() ExecuteInstancesBatchInfoRotationStrategyEnum {
	return ExecuteInstancesBatchInfoRotationStrategyEnum{
		CONTINUE: ExecuteInstancesBatchInfoRotationStrategy{
			value: "CONTINUE",
		},
		PAUSE: ExecuteInstancesBatchInfoRotationStrategy{
			value: "PAUSE",
		},
	}
}

func (c ExecuteInstancesBatchInfoRotationStrategy) Value() string {
	return c.value
}

func (c ExecuteInstancesBatchInfoRotationStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ExecuteInstancesBatchInfoRotationStrategy) UnmarshalJSON(b []byte) error {
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
