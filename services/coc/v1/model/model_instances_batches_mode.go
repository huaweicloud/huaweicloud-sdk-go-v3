package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// InstancesBatchesMode 需要分批的机器
type InstancesBatchesMode struct {

	// 分批策略：只支持自动分批
	BatchStrategy InstancesBatchesModeBatchStrategy `json:"batch_strategy"`

	// 目标主机实例
	TargetInstances []ResourceInstance `json:"target_instances"`
}

func (o InstancesBatchesMode) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstancesBatchesMode struct{}"
	}

	return strings.Join([]string{"InstancesBatchesMode", string(data)}, " ")
}

type InstancesBatchesModeBatchStrategy struct {
	value string
}

type InstancesBatchesModeBatchStrategyEnum struct {
	AUTO_BATCH InstancesBatchesModeBatchStrategy
}

func GetInstancesBatchesModeBatchStrategyEnum() InstancesBatchesModeBatchStrategyEnum {
	return InstancesBatchesModeBatchStrategyEnum{
		AUTO_BATCH: InstancesBatchesModeBatchStrategy{
			value: "AUTO_BATCH",
		},
	}
}

func (c InstancesBatchesModeBatchStrategy) Value() string {
	return c.value
}

func (c InstancesBatchesModeBatchStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstancesBatchesModeBatchStrategy) UnmarshalJSON(b []byte) error {
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
