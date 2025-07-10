package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmScheduleInstance 保存定时任务关联的实例选择条件或实例信息
type AlarmScheduleInstance struct {

	// 目标选择方式，枚举值：ALL 全部实例，MANUAL 手动选择 NONE
	TargetSelection AlarmScheduleInstanceTargetSelection `json:"target_selection"`

	// 实例信息
	TargetInstances *string `json:"target_instances,omitempty"`

	// 步骤号
	OrderNo int32 `json:"order_no"`

	// 实例分批策略 AUTO_BATCH 自动分批，MANUAL_BATCH手动分批，NONE不分批
	BatchStrategy *AlarmScheduleInstanceBatchStrategy `json:"batch_strategy,omitempty"`

	// 子目标实例
	SubTargetInstances *[]AlarmScheduleInstance `json:"sub_target_instances,omitempty"`
}

func (o AlarmScheduleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmScheduleInstance struct{}"
	}

	return strings.Join([]string{"AlarmScheduleInstance", string(data)}, " ")
}

type AlarmScheduleInstanceTargetSelection struct {
	value string
}

type AlarmScheduleInstanceTargetSelectionEnum struct {
	ALL    AlarmScheduleInstanceTargetSelection
	MANUAL AlarmScheduleInstanceTargetSelection
	NONE   AlarmScheduleInstanceTargetSelection
}

func GetAlarmScheduleInstanceTargetSelectionEnum() AlarmScheduleInstanceTargetSelectionEnum {
	return AlarmScheduleInstanceTargetSelectionEnum{
		ALL: AlarmScheduleInstanceTargetSelection{
			value: "ALL",
		},
		MANUAL: AlarmScheduleInstanceTargetSelection{
			value: "MANUAL",
		},
		NONE: AlarmScheduleInstanceTargetSelection{
			value: "NONE",
		},
	}
}

func (c AlarmScheduleInstanceTargetSelection) Value() string {
	return c.value
}

func (c AlarmScheduleInstanceTargetSelection) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmScheduleInstanceTargetSelection) UnmarshalJSON(b []byte) error {
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

type AlarmScheduleInstanceBatchStrategy struct {
	value string
}

type AlarmScheduleInstanceBatchStrategyEnum struct {
	AUTO_BATCH   AlarmScheduleInstanceBatchStrategy
	MANUAL_BATCH AlarmScheduleInstanceBatchStrategy
	NONE         AlarmScheduleInstanceBatchStrategy
}

func GetAlarmScheduleInstanceBatchStrategyEnum() AlarmScheduleInstanceBatchStrategyEnum {
	return AlarmScheduleInstanceBatchStrategyEnum{
		AUTO_BATCH: AlarmScheduleInstanceBatchStrategy{
			value: "AUTO_BATCH",
		},
		MANUAL_BATCH: AlarmScheduleInstanceBatchStrategy{
			value: "MANUAL_BATCH",
		},
		NONE: AlarmScheduleInstanceBatchStrategy{
			value: "NONE",
		},
	}
}

func (c AlarmScheduleInstanceBatchStrategy) Value() string {
	return c.value
}

func (c AlarmScheduleInstanceBatchStrategy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmScheduleInstanceBatchStrategy) UnmarshalJSON(b []byte) error {
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
