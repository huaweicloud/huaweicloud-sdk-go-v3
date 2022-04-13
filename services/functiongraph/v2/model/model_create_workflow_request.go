package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type CreateWorkflowRequest struct {
	Body *CreateWorkflowRequestBody `json:"body,omitempty"`
}

func (o CreateWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWorkflowRequest struct{}"
	}

	return strings.Join([]string{"CreateWorkflowRequest", string(data)}, " ")
}
