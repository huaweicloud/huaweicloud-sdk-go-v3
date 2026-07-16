package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowStepExecutionMetricsRequest Request Object
type ShowWorkflowStepExecutionMetricsRequest struct {

	// 工作流的ID。
	WorkflowId string `json:"workflow_id"`

	// 工作流执行ID。
	ExecutionId string `json:"execution_id"`

	// 工作流的一次执行中一个节点的执行ID。
	StepExecutionId string `json:"step_execution_id"`
}

func (o ShowWorkflowStepExecutionMetricsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowStepExecutionMetricsRequest struct{}"
	}

	return strings.Join([]string{"ShowWorkflowStepExecutionMetricsRequest", string(data)}, " ")
}
