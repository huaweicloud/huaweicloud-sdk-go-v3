package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// WorkflowTodo 工作流待办事项。
type WorkflowTodo struct {

	// 时间。
	Time *string `json:"time,omitempty"`

	// 运行时长。
	Duration *int32 `json:"duration,omitempty"`

	// Workflow工作流ID。
	WorkflowId *string `json:"workflow_id,omitempty"`

	// 工作流名称。填写1-64位，仅包含英文、数字、下划线（_）和中划线（-），并且以英文开头的名称。
	WorkflowName *string `json:"workflow_name,omitempty"`

	// 工作流执行ID。
	ExecutionId *string `json:"execution_id,omitempty"`

	// 节点名称。
	StepName *string `json:"step_name,omitempty"`

	// 节点的Title。
	StepTitle *string `json:"step_title,omitempty"`

	// 状态。
	Status *WorkflowTodoStatus `json:"status,omitempty"`
}

func (o WorkflowTodo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowTodo struct{}"
	}

	return strings.Join([]string{"WorkflowTodo", string(data)}, " ")
}

type WorkflowTodoStatus struct {
	value string
}

type WorkflowTodoStatusEnum struct {
	WAIT_INPUTS WorkflowTodoStatus
}

func GetWorkflowTodoStatusEnum() WorkflowTodoStatusEnum {
	return WorkflowTodoStatusEnum{
		WAIT_INPUTS: WorkflowTodoStatus{
			value: "wait_inputs",
		},
	}
}

func (c WorkflowTodoStatus) Value() string {
	return c.value
}

func (c WorkflowTodoStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *WorkflowTodoStatus) UnmarshalJSON(b []byte) error {
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
