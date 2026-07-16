package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkflowExecutionRequest Request Object
type UpdateWorkflowExecutionRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作流执行ID。
	ExecutionId string `json:"execution_id"`

	Body *WorkflowExecution `json:"body,omitempty"`
}

func (o UpdateWorkflowExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowExecutionRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowExecutionRequest", string(data)}, " ")
}
