package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteWorkflowsRequest Request Object
type BatchDeleteWorkflowsRequest struct {
	Body *WorkflowDeleteBody `json:"body,omitempty"`
}

func (o BatchDeleteWorkflowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteWorkflowsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteWorkflowsRequest", string(data)}, " ")
}
