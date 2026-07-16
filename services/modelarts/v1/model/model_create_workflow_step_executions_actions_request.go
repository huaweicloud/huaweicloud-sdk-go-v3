package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowStepExecutionsActionsRequest Request Object
type CreateWorkflowStepExecutionsActionsRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作流执行ID。
	ExecutionId string `json:"execution_id"`

	// 工作流的一次执行中一个节点的执行ID。
	StepExecutionId string `json:"step_execution_id"`

	Body *StepExecutionAction `json:"body,omitempty"`
}

func (o CreateWorkflowStepExecutionsActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowStepExecutionsActionsRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowStepExecutionsActionsRequest", string(data)}, " ")
}
