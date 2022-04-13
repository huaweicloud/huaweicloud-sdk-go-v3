package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type CreateTasksRequestBody struct {
	// 任务名称

	TaskName string `json:"task_name"`
	// 待扫描的目标网址

	Url string `json:"url"`
	// 扫描任务类型:   * normal - 普通任务   * monitor - 监测任务

	TaskType *CreateTasksRequestBodyTaskType `json:"task_type,omitempty"`
	// 普通任务的定时启动时间

	Timer *string `json:"timer,omitempty"`
	// 监测任务的定时触发时间

	TriggerTime *string `json:"trigger_time,omitempty"`
	// 监测任务的定时触发周期:   * everyday - 每日   * threedays - 每三天   * everyweek - 每星期   * everymonth - 每月

	TaskPeriod *CreateTasksRequestBodyTaskPeriod `json:"task_period,omitempty"`

	TaskConfig *TaskSettingsTaskConfig `json:"task_config,omitempty"`
}

func (o CreateTasksRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateTasksRequestBody struct{}"
	}

	return strings.Join([]string{"CreateTasksRequestBody", string(data)}, " ")
}

type CreateTasksRequestBodyTaskType struct {
	value string
}

type CreateTasksRequestBodyTaskTypeEnum struct {
	NORMAL  CreateTasksRequestBodyTaskType
	MONITOR CreateTasksRequestBodyTaskType
}

func GetCreateTasksRequestBodyTaskTypeEnum() CreateTasksRequestBodyTaskTypeEnum {
	return CreateTasksRequestBodyTaskTypeEnum{
		NORMAL: CreateTasksRequestBodyTaskType{
			value: "normal",
		},
		MONITOR: CreateTasksRequestBodyTaskType{
			value: "monitor",
		},
	}
}

func (c CreateTasksRequestBodyTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTasksRequestBodyTaskType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}

type CreateTasksRequestBodyTaskPeriod struct {
	value string
}

type CreateTasksRequestBodyTaskPeriodEnum struct {
	EVERYDAY   CreateTasksRequestBodyTaskPeriod
	THREEDAYS  CreateTasksRequestBodyTaskPeriod
	EVERYWEEK  CreateTasksRequestBodyTaskPeriod
	EVERYMONTH CreateTasksRequestBodyTaskPeriod
}

func GetCreateTasksRequestBodyTaskPeriodEnum() CreateTasksRequestBodyTaskPeriodEnum {
	return CreateTasksRequestBodyTaskPeriodEnum{
		EVERYDAY: CreateTasksRequestBodyTaskPeriod{
			value: "everyday",
		},
		THREEDAYS: CreateTasksRequestBodyTaskPeriod{
			value: "threedays",
		},
		EVERYWEEK: CreateTasksRequestBodyTaskPeriod{
			value: "everyweek",
		},
		EVERYMONTH: CreateTasksRequestBodyTaskPeriod{
			value: "everymonth",
		},
	}
}

func (c CreateTasksRequestBodyTaskPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateTasksRequestBodyTaskPeriod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
