package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWorkflowServiceAuthRequest Request Object
type CreateWorkflowServiceAuthRequest struct {
	Body *WorkflowMainServiceAuthReq `json:"body,omitempty"`
}

func (o CreateWorkflowServiceAuthRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowServiceAuthRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowServiceAuthRequest", string(data)}, " ")
}
