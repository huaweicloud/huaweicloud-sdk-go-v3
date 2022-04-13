package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StartWorkflowExecutionRequest struct {
	// 函数工作流ID

	WorkflowId string `json:"workflow_id"`
	// workflowRun task create time

	XCreateTime *string `json:"X-Create-Time,omitempty"`
	// workflowRun id

	XWorkflowRunID *string `json:"X-WorkflowRun-ID,omitempty"`

	Body *StartWorkflowExecutionRequestBody `json:"body,omitempty"`
}

func (o StartWorkflowExecutionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StartWorkflowExecutionRequest struct{}"
	}

	return strings.Join([]string{"StartWorkflowExecutionRequest", string(data)}, " ")
}
