package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type RetryWorkFlowRequest struct {
	// 函数工作流ID

	WorkflowId string `json:"workflow_id"`
	// 函数流执行实例ID

	ExecutionId string `json:"execution_id"`
}

func (o RetryWorkFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RetryWorkFlowRequest struct{}"
	}

	return strings.Join([]string{"RetryWorkFlowRequest", string(data)}, " ")
}
