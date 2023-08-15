package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWorkFlowRequest Request Object
type UpdateWorkFlowRequest struct {

	// 函数工作流ID
	WorkflowId string `json:"workflow_id"`

	Body *WorkflowCreateBody `json:"body,omitempty"`
}

func (o UpdateWorkFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWorkFlowRequest struct{}"
	}

	return strings.Join([]string{"UpdateWorkFlowRequest", string(data)}, " ")
}
