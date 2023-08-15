package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowRequest Request Object
type CreateWorkflowRequest struct {
	Body *WorkflowRequestBody `json:"body,omitempty"`
}

func (o CreateWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowRequest", string(data)}, " ")
}
