package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowExecutionsActionsRequest Request Object
type CreateWorkflowExecutionsActionsRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作流执行ID。
	ExecutionId string `json:"execution_id"`

	Body *ExecutionAction `json:"body,omitempty"`
}

func (o CreateWorkflowExecutionsActionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowExecutionsActionsRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowExecutionsActionsRequest", string(data)}, " ")
}
