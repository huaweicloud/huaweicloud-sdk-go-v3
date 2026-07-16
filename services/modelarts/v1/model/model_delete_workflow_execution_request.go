package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteWorkflowExecutionRequest Request Object
type DeleteWorkflowExecutionRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作流执行ID。
	ExecutionId string `json:"execution_id"`
}

func (o DeleteWorkflowExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteWorkflowExecutionRequest struct{}"
	}

	return strings.Join([]string{"DeleteWorkflowExecutionRequest", string(data)}, " ")
}
