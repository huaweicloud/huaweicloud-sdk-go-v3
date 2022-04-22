package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowTasksResponse struct {

	// 任务名称
	TaskName string `json:"task_name"`

	// 待扫描的目标网址
	Url string `json:"url"`

	// 扫描任务类型:   * normal - 普通任务   * monitor - 监测任务
	TaskType *ShowTasksResponseTaskType `json:"task_type,omitempty"`

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
	TaskStatus *ShowTasksResponseTaskStatus `json:"task_status,omitempty"`

	// 监测任务状态:   * running - 正在运行   * waiting - 正在等待   * finished - 已完成
	ScheduleStatus *ShowTasksResponseScheduleStatus `json:"schedule_status,omitempty"`

	// 任务进度
	Progress *int32 `json:"progress,omitempty"`

	// 任务状态描述
	Reason *string `json:"reason,omitempty"`

	// 包总数
	PackNum *int64 `json:"pack_num,omitempty"`

	// 安全分数
	Score *int32 `json:"score,omitempty"`

	// 安全等级:   * - safety  : 安全   * - average : 中风险   * - highrisk: 高风险
	SafeLevel *ShowTasksResponseSafeLevel `json:"safe_level,omitempty"`

	Statistics     *VulnsLevel `json:"statistics,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o ShowTasksResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTasksResponse struct{}"
	}

	return strings.Join([]string{"ShowTasksResponse", string(data)}, " ")
}

type ShowTasksResponseTaskType struct {
	value string
}

type ShowTasksResponseTaskTypeEnum struct {
	NORMAL  ShowTasksResponseTaskType
	MONITOR ShowTasksResponseTaskType
}

func GetShowTasksResponseTaskTypeEnum() ShowTasksResponseTaskTypeEnum {
	return ShowTasksResponseTaskTypeEnum{
		NORMAL: ShowTasksResponseTaskType{
			value: "normal",
		},
		MONITOR: ShowTasksResponseTaskType{
			value: "monitor",
		},
	}
}

func (c ShowTasksResponseTaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTasksResponseTaskType) UnmarshalJSON(b []byte) error {
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

type ShowTasksResponseTaskStatus struct {
	value string
}

type ShowTasksResponseTaskStatusEnum struct {
	RUNNING  ShowTasksResponseTaskStatus
	SUCCESS  ShowTasksResponseTaskStatus
	CANCELED ShowTasksResponseTaskStatus
	WAITING  ShowTasksResponseTaskStatus
	FAILURE  ShowTasksResponseTaskStatus
}

func GetShowTasksResponseTaskStatusEnum() ShowTasksResponseTaskStatusEnum {
	return ShowTasksResponseTaskStatusEnum{
		RUNNING: ShowTasksResponseTaskStatus{
			value: "running",
		},
		SUCCESS: ShowTasksResponseTaskStatus{
			value: "success",
		},
		CANCELED: ShowTasksResponseTaskStatus{
			value: "canceled",
		},
		WAITING: ShowTasksResponseTaskStatus{
			value: "waiting",
		},
		FAILURE: ShowTasksResponseTaskStatus{
			value: "failure",
		},
	}
}

func (c ShowTasksResponseTaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTasksResponseTaskStatus) UnmarshalJSON(b []byte) error {
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

type ShowTasksResponseScheduleStatus struct {
	value string
}

type ShowTasksResponseScheduleStatusEnum struct {
	RUNNING  ShowTasksResponseScheduleStatus
	WAITING  ShowTasksResponseScheduleStatus
	FINISHED ShowTasksResponseScheduleStatus
}

func GetShowTasksResponseScheduleStatusEnum() ShowTasksResponseScheduleStatusEnum {
	return ShowTasksResponseScheduleStatusEnum{
		RUNNING: ShowTasksResponseScheduleStatus{
			value: "running",
		},
		WAITING: ShowTasksResponseScheduleStatus{
			value: "waiting",
		},
		FINISHED: ShowTasksResponseScheduleStatus{
			value: "finished",
		},
	}
}

func (c ShowTasksResponseScheduleStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTasksResponseScheduleStatus) UnmarshalJSON(b []byte) error {
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

type ShowTasksResponseSafeLevel struct {
	value string
}

type ShowTasksResponseSafeLevelEnum struct {
	SAFETY   ShowTasksResponseSafeLevel
	AVERAGE  ShowTasksResponseSafeLevel
	HIGHRISK ShowTasksResponseSafeLevel
}

func GetShowTasksResponseSafeLevelEnum() ShowTasksResponseSafeLevelEnum {
	return ShowTasksResponseSafeLevelEnum{
		SAFETY: ShowTasksResponseSafeLevel{
			value: "safety",
		},
		AVERAGE: ShowTasksResponseSafeLevel{
			value: "average",
		},
		HIGHRISK: ShowTasksResponseSafeLevel{
			value: "highrisk",
		},
	}
}

func (c ShowTasksResponseSafeLevel) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowTasksResponseSafeLevel) UnmarshalJSON(b []byte) error {
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
