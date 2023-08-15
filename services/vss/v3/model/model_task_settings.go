package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type TaskSettings struct {

	// 普通任务的定时启动时间
	Timer *string `json:"timer,omitempty"`

	// 监测任务的定时触发时间
	TriggerTime *string `json:"trigger_time,omitempty"`

	// 监测任务的定时触发周期:   * everyday - 每日   * threedays - 每三天   * everyweek - 每星期   * everymonth - 每月
	TaskPeriod *TaskSettingsTaskPeriod `json:"task_period,omitempty"`

	TaskConfig *TaskSettingsTaskConfig `json:"task_config,omitempty"`
}

func (o TaskSettings) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskSettings struct{}"
	}

	return strings.Join([]string{"TaskSettings", string(data)}, " ")
}

type TaskSettingsTaskPeriod struct {
	value string
}

type TaskSettingsTaskPeriodEnum struct {
	EVERYDAY   TaskSettingsTaskPeriod
	THREEDAYS  TaskSettingsTaskPeriod
	EVERYWEEK  TaskSettingsTaskPeriod
	EVERYMONTH TaskSettingsTaskPeriod
}

func GetTaskSettingsTaskPeriodEnum() TaskSettingsTaskPeriodEnum {
	return TaskSettingsTaskPeriodEnum{
		EVERYDAY: TaskSettingsTaskPeriod{
			value: "everyday",
		},
		THREEDAYS: TaskSettingsTaskPeriod{
			value: "threedays",
		},
		EVERYWEEK: TaskSettingsTaskPeriod{
			value: "everyweek",
		},
		EVERYMONTH: TaskSettingsTaskPeriod{
			value: "everymonth",
		},
	}
}

func (c TaskSettingsTaskPeriod) Value() string {
	return c.value
}

func (c TaskSettingsTaskPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskSettingsTaskPeriod) UnmarshalJSON(b []byte) error {
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
