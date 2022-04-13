package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type StopWorkFlowRequest struct {
	// 函数工作流ID

	WorkflowId string `json:"workflow_id"`
	// 函数流执行实例ID

	ExecutionId string `json:"execution_id"`
}

func (o StopWorkFlowRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopWorkFlowRequest struct{}"
	}

	return strings.Join([]string{"StopWorkFlowRequest", string(data)}, " ")
}
