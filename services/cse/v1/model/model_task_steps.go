package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// TaskSteps 处理阶段
type TaskSteps struct {

	// 处理阶段名称
	TaskName *string `json:"task_name,omitempty"`

	// 当前处理阶段包含的处理步骤名称列表
	TaskNames *[]string `json:"task_names,omitempty"`

	// 处理阶段状态
	Status *TaskStepsStatus `json:"status,omitempty"`

	// 处理阶段开始时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 处理阶段结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	TaskExecutorBrief *TaskExecutorBrief `json:"task_executor_brief,omitempty"`

	// 处理步骤
	Tasks *[]Task `json:"tasks,omitempty"`
}

func (o TaskSteps) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskSteps struct{}"
	}

	return strings.Join([]string{"TaskSteps", string(data)}, " ")
}

type TaskStepsStatus struct {
	value string
}

type TaskStepsStatusEnum struct {
	INIT      TaskStepsStatus
	EXECUTING TaskStepsStatus
	ERROR     TaskStepsStatus
	TIMEOUT   TaskStepsStatus
	FINISHED  TaskStepsStatus
}

func GetTaskStepsStatusEnum() TaskStepsStatusEnum {
	return TaskStepsStatusEnum{
		INIT: TaskStepsStatus{
			value: "Init",
		},
		EXECUTING: TaskStepsStatus{
			value: "Executing",
		},
		ERROR: TaskStepsStatus{
			value: "Error",
		},
		TIMEOUT: TaskStepsStatus{
			value: "Timeout",
		},
		FINISHED: TaskStepsStatus{
			value: "Finished",
		},
	}
}

func (c TaskStepsStatus) Value() string {
	return c.value
}

func (c TaskStepsStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskStepsStatus) UnmarshalJSON(b []byte) error {
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
