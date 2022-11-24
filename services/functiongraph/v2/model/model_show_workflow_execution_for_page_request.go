package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowWorkflowExecutionForPageRequest struct {

	// 函数工作流ID
	WorkflowId string `json:"workflow_id"`

	Body *QueryRunListParam `json:"body,omitempty"`
}

func (o ShowWorkflowExecutionForPageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowExecutionForPageRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowExecutionForPageRequest", string(data)}, " ")
}
