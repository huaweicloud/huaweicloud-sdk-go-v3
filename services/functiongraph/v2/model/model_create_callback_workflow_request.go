package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCallbackWorkflowRequest Request Object
type CreateCallbackWorkflowRequest struct {

	// 函数工作流ID
	WorkflowId string `json:"workflow_id"`

	// workflow run id
	XWorkflowRunId string `json:"X-Workflow-Run-Id"`

	// workflow state id
	XWorkflowStateId string `json:"X-Workflow-State-Id"`

	Body *CallbackWorkflowRequestBody `json:"body,omitempty"`
}

func (o CreateCallbackWorkflowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCallbackWorkflowRequest struct{}"
	}

	return strings.Join([]string{"CreateCallbackWorkflowRequest", string(data)}, " ")
}
