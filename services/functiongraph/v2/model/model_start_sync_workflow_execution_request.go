package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartSyncWorkflowExecutionRequest Request Object
type StartSyncWorkflowExecutionRequest struct {

	// 函数流定义ID
	WorkflowId string `json:"workflow_id"`

	Body *FlowExecuteBody `json:"body,omitempty"`
}

func (o StartSyncWorkflowExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartSyncWorkflowExecutionRequest struct{}"
	}

	return strings.Join([]string{"StartSyncWorkflowExecutionRequest", string(data)}, " ")
}
