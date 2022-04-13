package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchDeleteWorkflowsRequest struct {
	Body *BatchDeleteWorkflowsRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteWorkflowsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteWorkflowsRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteWorkflowsRequest", string(data)}, " ")
}
