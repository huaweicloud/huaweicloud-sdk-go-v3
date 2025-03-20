package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ScheduleInstance 保存定时任务关联的实例选择条件或实例信息
type ScheduleInstance struct {

	// 目标选择方式，枚举值：ALL 全部实例，MANUAL 手动选择, NONE
	TargetSelection ScheduleInstanceTargetSelection `json:"target_selection"`

	TargetResource *TargetResource `json:"target_resource,omitempty"`

	// 实例信息
	TargetInstances *string `json:"target_instances,omitempty"`

	// 步骤号
	OrderNo int32 `json:"order_no"`

	// 实例分批策略(AUTO_BATCH,MANUAL_BATCH,NONE)
	BatchStrategy *ScheduleInstanceBatchStrategy `json:"batch_strategy,omitempty"`

	// 目标实例
	SubTargetInstances *[]ScheduleInstance `json:"sub_target_instances,omitempty"`
}

func (o ScheduleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScheduleInstance struct{}"
	}

	return strings.Join([]string{"ScheduleInstance", string(data)}, " ")
}

type ScheduleInstanceTargetSelection struct {
	value string
}

type ScheduleInstanceTargetSelectionEnum struct {
	ALL    ScheduleInstanceTargetSelection
	MANUAL ScheduleInstanceTargetSelection
	NONE   ScheduleInstanceTargetSelection
}

func GetScheduleInstanceTargetSelectionEnum() ScheduleInstanceTargetSelectionEnum {
	return ScheduleInstanceTargetSelectionEnum{
		ALL: ScheduleInstanceTargetSelection{
			value: "ALL",
		},
		MANUAL: ScheduleInstanceTargetSelection{
			value: "MANUAL",
		},
		NONE: ScheduleInstanceTargetSelection{
			value: "NONE",
		},
	}
}

func (c ScheduleInstanceTargetSelection) Value() string {
	return c.value
}

func (c ScheduleInstanceTargetSelection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduleInstanceTargetSelection) UnmarshalJSON(b []byte) error {
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

type ScheduleInstanceBatchStrategy struct {
	value string
}

type ScheduleInstanceBatchStrategyEnum struct {
	AUTO_BATCH   ScheduleInstanceBatchStrategy
	MANUAL_BATCH ScheduleInstanceBatchStrategy
	NONE         ScheduleInstanceBatchStrategy
}

func GetScheduleInstanceBatchStrategyEnum() ScheduleInstanceBatchStrategyEnum {
	return ScheduleInstanceBatchStrategyEnum{
		AUTO_BATCH: ScheduleInstanceBatchStrategy{
			value: "AUTO_BATCH",
		},
		MANUAL_BATCH: ScheduleInstanceBatchStrategy{
			value: "MANUAL_BATCH",
		},
		NONE: ScheduleInstanceBatchStrategy{
			value: "NONE",
		},
	}
}

func (c ScheduleInstanceBatchStrategy) Value() string {
	return c.value
}

func (c ScheduleInstanceBatchStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ScheduleInstanceBatchStrategy) UnmarshalJSON(b []byte) error {
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
