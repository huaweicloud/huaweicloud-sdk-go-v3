package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowTasksResponseBody struct {
	// 任务名称

	TaskName string `json:"task_name"`
	// 待扫描的目标网址

	Url string `json:"url"`
	// 扫描任务类型:   * normal - 普通任务   * monitor - 监测任务

	TaskType *ShowTasksResponseBodyTaskType `json:"task_type,omitempty"`
	// 任务ID

	TaskId *string `json:"task_id,omitempty"`
	// 域名

	DomainName *string `json:"domain_name,omitempty"`

	TaskSettings *TaskSettings `json:"task_settings,omitempty"`
	// 创建任务的时间

	CreateTime *string `json:"create_time,omitempty"`
	// 任务启动的时间

	StartTime *string `json:"start_time,omitempty"`
	// 任务结束的时间

	EndTime *string `json:"end_time,omitempty"`
	// 任务状态:   * running - 正在运行   * success - 成功   * canceled - 已取消   * waiting - 正在等待   * failure - 失败

	TaskStatus *ShowTasksResponseBodyTaskStatus `json:"task_status,omitempty"`
	// 监测任务状态:   * running - 正在运行   * waiting - 正在等待   * finished - 已完成

	ScheduleStatus *ShowTasksResponseBodyScheduleStatus `json:"schedule_status,omitempty"`
	// 任务进度

	Progress *int32 `json:"progress,omitempty"`
	// 任务状态描述

	Reason *string `json:"reason,omitempty"`
	// 包总数

	PackNum *int64 `json:"pack_num,omitempty"`
	// 安全分数

	Score *int32 `json:"score,omitempty"`
	// 安全等级:   * - safety  : 安全   * - average : 中风险   * - highrisk: 高风险

	SafeLevel *ShowTasksResponseBodySafeLevel `json:"safe_level,omitempty"`

	Statistics *VulnsLevel `json:"statistics,omitempty"`
}

func (o ShowTasksResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTasksResponseBody struct{}"
	}

	return strings.Join([]string{"ShowTasksResponseBody", string(data)}, " ")
}

type ShowTasksResponseBodyTaskType struct {
	value string
}

type ShowTasksResponseBodyTaskTypeEnum struct {
	NORMAL  ShowTasksResponseBodyTaskType
	MONITOR ShowTasksResponseBodyTaskType
}

func GetShowTasksResponseBodyTaskTypeEnum() ShowTasksResponseBodyTaskTypeEnum {
	return ShowTasksResponseBodyTaskTypeEnum{
		NORMAL: ShowTasksResponseBodyTaskType{
			value: "normal",
		},
		MONITOR: ShowTasksResponseBodyTaskType{
			value: "monitor",
		},
	}
}

func (c ShowTasksResponseBodyTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTasksResponseBodyTaskType) UnmarshalJSON(b []byte) error {
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

type ShowTasksResponseBodyTaskStatus struct {
	value string
}

type ShowTasksResponseBodyTaskStatusEnum struct {
	RUNNING  ShowTasksResponseBodyTaskStatus
	SUCCESS  ShowTasksResponseBodyTaskStatus
	CANCELED ShowTasksResponseBodyTaskStatus
	WAITING  ShowTasksResponseBodyTaskStatus
	FAILURE  ShowTasksResponseBodyTaskStatus
}

func GetShowTasksResponseBodyTaskStatusEnum() ShowTasksResponseBodyTaskStatusEnum {
	return ShowTasksResponseBodyTaskStatusEnum{
		RUNNING: ShowTasksResponseBodyTaskStatus{
			value: "running",
		},
		SUCCESS: ShowTasksResponseBodyTaskStatus{
			value: "success",
		},
		CANCELED: ShowTasksResponseBodyTaskStatus{
			value: "canceled",
		},
		WAITING: ShowTasksResponseBodyTaskStatus{
			value: "waiting",
		},
		FAILURE: ShowTasksResponseBodyTaskStatus{
			value: "failure",
		},
	}
}

func (c ShowTasksResponseBodyTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTasksResponseBodyTaskStatus) UnmarshalJSON(b []byte) error {
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

type ShowTasksResponseBodyScheduleStatus struct {
	value string
}

type ShowTasksResponseBodyScheduleStatusEnum struct {
	RUNNING  ShowTasksResponseBodyScheduleStatus
	WAITING  ShowTasksResponseBodyScheduleStatus
	FINISHED ShowTasksResponseBodyScheduleStatus
}

func GetShowTasksResponseBodyScheduleStatusEnum() ShowTasksResponseBodyScheduleStatusEnum {
	return ShowTasksResponseBodyScheduleStatusEnum{
		RUNNING: ShowTasksResponseBodyScheduleStatus{
			value: "running",
		},
		WAITING: ShowTasksResponseBodyScheduleStatus{
			value: "waiting",
		},
		FINISHED: ShowTasksResponseBodyScheduleStatus{
			value: "finished",
		},
	}
}

func (c ShowTasksResponseBodyScheduleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTasksResponseBodyScheduleStatus) UnmarshalJSON(b []byte) error {
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

type ShowTasksResponseBodySafeLevel struct {
	value string
}

type ShowTasksResponseBodySafeLevelEnum struct {
	SAFETY   ShowTasksResponseBodySafeLevel
	AVERAGE  ShowTasksResponseBodySafeLevel
	HIGHRISK ShowTasksResponseBodySafeLevel
}

func GetShowTasksResponseBodySafeLevelEnum() ShowTasksResponseBodySafeLevelEnum {
	return ShowTasksResponseBodySafeLevelEnum{
		SAFETY: ShowTasksResponseBodySafeLevel{
			value: "safety",
		},
		AVERAGE: ShowTasksResponseBodySafeLevel{
			value: "average",
		},
		HIGHRISK: ShowTasksResponseBodySafeLevel{
			value: "highrisk",
		},
	}
}

func (c ShowTasksResponseBodySafeLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTasksResponseBodySafeLevel) UnmarshalJSON(b []byte) error {
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
