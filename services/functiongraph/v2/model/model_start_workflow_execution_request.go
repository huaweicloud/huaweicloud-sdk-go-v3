package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StartWorkflowExecutionRequest Request Object
type StartWorkflowExecutionRequest struct {

	// 函数流定义ID
	WorkflowId string `json:"workflow_id"`

	// workflowRun task create time
	XCreateTime *string `json:"X-Create-Time,omitempty"`

	// workflowRun id
	XWorkflowRunID *string `json:"X-WorkflowRun-ID,omitempty"`

	// Combines the output of the previous node with the input of the next node into an input.
	XWorkflowRunMergeFnParameters *string `json:"X-WorkflowRun-MergeFnParameters,omitempty"`

	Body *FlowExecuteBody `json:"body,omitempty"`
}

func (o StartWorkflowExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartWorkflowExecutionRequest struct{}"
	}

	return strings.Join([]string{"StartWorkflowExecutionRequest", string(data)}, " ")
}
