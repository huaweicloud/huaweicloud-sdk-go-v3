package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Task 子任务
type Task struct {

	// 子任务所属任务ID
	JobId *int32 `json:"jobId,omitempty"`

	// 子任务ID，使用uuid
	Id *int64 `json:"id,omitempty"`

	// 子任务的类型
	Type *TaskType `json:"type,omitempty"`

	// 子任务的执行者
	Assigned *string `json:"assigned,omitempty"`

	// 子任务名称
	TaskName *string `json:"taskName,omitempty"`

	// 子任务所属引擎名称
	EngineName *string `json:"engineName,omitempty"`

	// 子任务执行的顺序, 从小到大
	TaskOrder *int32 `json:"taskOrder,omitempty"`

	// 子任务状态
	Status *TaskStatus `json:"status,omitempty"`

	// 子任务开始时间
	StartTime *int64 `json:"startTime,omitempty"`

	// 子任务结束时间
	EndTime *int64 `json:"endTime,omitempty"`

	// 子任务创建时间
	CreateTime *int64 `json:"createTime,omitempty"`

	// 子任务更新时间
	UpdateTime *int64 `json:"updateTime,omitempty"`

	// 子任务是否超时
	Timeout *int32 `json:"timeout,omitempty"`

	// 子任务详细信息，执行过程中产生的辅助信息
	Log *string `json:"log,omitempty"`

	// 子任务输出信息
	Output *string `json:"output,omitempty"`

	TaskExecutorBrief *TaskExecutorBrief `json:"taskExecutorBrief,omitempty"`
}

func (o Task) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Task struct{}"
	}

	return strings.Join([]string{"Task", string(data)}, " ")
}

type TaskType struct {
	value string
}

type TaskTypeEnum struct {
	CREATE  TaskType
	DELETE  TaskType
	UPGRADE TaskType
	MODIFY  TaskType
}

func GetTaskTypeEnum() TaskTypeEnum {
	return TaskTypeEnum{
		CREATE: TaskType{
			value: "Create",
		},
		DELETE: TaskType{
			value: "Delete",
		},
		UPGRADE: TaskType{
			value: "Upgrade",
		},
		MODIFY: TaskType{
			value: "Modify",
		},
	}
}

func (c TaskType) Value() string {
	return c.value
}

func (c TaskType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskType) UnmarshalJSON(b []byte) error {
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

type TaskStatus struct {
	value string
}

type TaskStatusEnum struct {
	INIT      TaskStatus
	EXECUTING TaskStatus
	ERROR     TaskStatus
	TIMEOUT   TaskStatus
	FINISHED  TaskStatus
}

func GetTaskStatusEnum() TaskStatusEnum {
	return TaskStatusEnum{
		INIT: TaskStatus{
			value: "Init",
		},
		EXECUTING: TaskStatus{
			value: "Executing",
		},
		ERROR: TaskStatus{
			value: "Error",
		},
		TIMEOUT: TaskStatus{
			value: "Timeout",
		},
		FINISHED: TaskStatus{
			value: "Finished",
		},
	}
}

func (c TaskStatus) Value() string {
	return c.value
}

func (c TaskStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *TaskStatus) UnmarshalJSON(b []byte) error {
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
