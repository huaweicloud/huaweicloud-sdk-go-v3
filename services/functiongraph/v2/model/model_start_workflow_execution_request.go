package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartWorkflowExecutionRequest struct {

	// 函数流定义ID
	WorkflowId string `json:"workflow_id" xml:"workflow_id"`

	// workflowRun task create time
	XCreateTime *string `json:"X-Create-Time,omitempty" xml:"X-Create-Time"`

	// workflowRun id
	XWorkflowRunID *string `json:"X-WorkflowRun-ID,omitempty" xml:"X-WorkflowRun-ID"`

	// Combines the output of the previous node with the input of the next node into an input.
	XWorkflowRunMergeFnParameters *string `json:"X-WorkflowRun-MergeFnParameters,omitempty" xml:"X-WorkflowRun-MergeFnParameters"`

	Body *FlowExecuteBody `json:"body,omitempty" xml:"body"`
}

func (o StartWorkflowExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartWorkflowExecutionRequest struct{}"
	}

	return strings.Join([]string{"StartWorkflowExecutionRequest", string(data)}, " ")
}
