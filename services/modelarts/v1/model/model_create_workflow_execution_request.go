package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowExecutionRequest Request Object
type CreateWorkflowExecutionRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	Body *WorkflowExecution `json:"body,omitempty"`
}

func (o CreateWorkflowExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowExecutionRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowExecutionRequest", string(data)}, " ")
}
