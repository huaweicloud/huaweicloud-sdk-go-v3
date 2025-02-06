package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskListItemVo 单个任务详情数据
type TaskListItemVo struct {

	// 任务ID
	TaskId string `json:"task_id"`

	// 任务状态
	Status TaskListItemVoStatus `json:"status"`

	// 开始时间
	StartTime string `json:"start_time"`

	// 结束时间
	EndTime string `json:"end_time"`

	// 创建时间
	CreateTime string `json:"create_time"`
}

func (o TaskListItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskListItemVo struct{}"
	}

	return strings.Join([]string{"TaskListItemVo", string(data)}, " ")
}

type TaskListItemVoStatus struct {
	value string
}

type TaskListItemVoStatusEnum struct {
	RUNNING   TaskListItemVoStatus
	PENDING   TaskListItemVoStatus
	EXPIRED   TaskListItemVoStatus
	UNKNOWN   TaskListItemVoStatus
	FAILED    TaskListItemVoStatus
	SUCCEEDED TaskListItemVoStatus
	STOPPED   TaskListItemVoStatus
	DELETED   TaskListItemVoStatus
}

func GetTaskListItemVoStatusEnum() TaskListItemVoStatusEnum {
	return TaskListItemVoStatusEnum{
		RUNNING: TaskListItemVoStatus{
			value: "Running",
		},
		PENDING: TaskListItemVoStatus{
			value: "Pending",
		},
		EXPIRED: TaskListItemVoStatus{
			value: "Expired",
		},
		UNKNOWN: TaskListItemVoStatus{
			value: "Unknown",
		},
		FAILED: TaskListItemVoStatus{
			value: "Failed",
		},
		SUCCEEDED: TaskListItemVoStatus{
			value: "Succeeded",
		},
		STOPPED: TaskListItemVoStatus{
			value: "Stopped",
		},
		DELETED: TaskListItemVoStatus{
			value: "Deleted",
		},
	}
}

func (c TaskListItemVoStatus) Value() string {
	return c.value
}

func (c TaskListItemVoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskListItemVoStatus) UnmarshalJSON(b []byte) error {
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
