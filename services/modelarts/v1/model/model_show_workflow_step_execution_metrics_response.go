package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWorkflowStepExecutionMetricsResponse Response Object
type ShowWorkflowStepExecutionMetricsResponse struct {
	Body           *[]WorkflowStepMetric `json:"body,omitempty"`
	HttpStatusCode int                   `json:"-"`
}

func (o ShowWorkflowStepExecutionMetricsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWorkflowStepExecutionMetricsResponse struct{}"
	}

	return strings.Join([]string{"ShowWorkflowStepExecutionMetricsResponse", string(data)}, " ")
}
