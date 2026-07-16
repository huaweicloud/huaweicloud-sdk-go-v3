package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowExecutionRequest Request Object
type ShowWorkflowExecutionRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作流执行ID。
	ExecutionId string `json:"execution_id"`
}

func (o ShowWorkflowExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowExecutionRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowExecutionRequest", string(data)}, " ")
}
