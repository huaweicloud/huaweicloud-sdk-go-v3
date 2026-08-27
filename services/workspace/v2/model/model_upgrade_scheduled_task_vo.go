package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpgradeScheduledTaskVo 升级任务VO
type UpgradeScheduledTaskVo struct {

	// 任务ID
	Id *string `json:"id,omitempty"`

	// 任务名称
	TaskName *string `json:"task_name,omitempty"`

	// 任务类型：0-云桌面 1-应用服务器 2-镜像
	TaskType *int32 `json:"task_type,omitempty"`

	// 执行周期类型：FIXED_TIME-指定时间 DAY-按天 WEEK-按周 MONTH-按月
	ScheduledType *UpgradeScheduledTaskVoScheduledType `json:"scheduled_type,omitempty"`

	// 时区
	Timezone *string `json:"timezone,omitempty"`

	// 最近一次执行情况：SUCCESS-成功 FAILED-失败 RUNNING-执行中 WAITING-等待
	LastExecuteStatus *string `json:"last_execute_status,omitempty"`

	// 下次执行时间
	NextExecuteTime *string `json:"next_execute_time,omitempty"`

	// 启用状态：0-未启用 1-启用
	IsEnable *int32 `json:"is_enable,omitempty"`

	// 目标版本
	TargetVersion *string `json:"target_version,omitempty"`

	// 执行策略：0-全量下发 1-灰度下发
	ExecuteStrategy *int32 `json:"execute_strategy,omitempty"`

	// 任务描述
	Description *string `json:"description,omitempty"`
}

func (o UpgradeScheduledTaskVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpgradeScheduledTaskVo struct{}"
	}

	return strings.Join([]string{"UpgradeScheduledTaskVo", string(data)}, " ")
}

type UpgradeScheduledTaskVoScheduledType struct {
	value string
}

type UpgradeScheduledTaskVoScheduledTypeEnum struct {
	FIXED_TIME UpgradeScheduledTaskVoScheduledType
	DAY        UpgradeScheduledTaskVoScheduledType
	WEEK       UpgradeScheduledTaskVoScheduledType
	MONTH      UpgradeScheduledTaskVoScheduledType
}

func GetUpgradeScheduledTaskVoScheduledTypeEnum() UpgradeScheduledTaskVoScheduledTypeEnum {
	return UpgradeScheduledTaskVoScheduledTypeEnum{
		FIXED_TIME: UpgradeScheduledTaskVoScheduledType{
			value: "FIXED_TIME",
		},
		DAY: UpgradeScheduledTaskVoScheduledType{
			value: "DAY",
		},
		WEEK: UpgradeScheduledTaskVoScheduledType{
			value: "WEEK",
		},
		MONTH: UpgradeScheduledTaskVoScheduledType{
			value: "MONTH",
		},
	}
}

func (c UpgradeScheduledTaskVoScheduledType) Value() string {
	return c.value
}

func (c UpgradeScheduledTaskVoScheduledType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpgradeScheduledTaskVoScheduledType) UnmarshalJSON(b []byte) error {
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
