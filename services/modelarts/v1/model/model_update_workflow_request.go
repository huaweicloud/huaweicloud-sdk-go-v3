package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkflowRequest Request Object
type UpdateWorkflowRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	Body *WorkflowUpdate `json:"body,omitempty"`
}

func (o UpdateWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkflowRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkflowRequest", string(data)}, " ")
}
