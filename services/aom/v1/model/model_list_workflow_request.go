package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListWorkflowRequest struct {
	Body *WorkflowQueryParam `json:"body,omitempty"`
}

func (o ListWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowRequest struct{}"
	}

	return strings.Join([]string{"ListWorkflowRequest", string(data)}, " ")
}
