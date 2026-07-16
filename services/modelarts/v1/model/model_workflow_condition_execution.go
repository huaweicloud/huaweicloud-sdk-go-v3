package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowConditionExecution workflow condition execution
type WorkflowConditionExecution struct {

	// 执行结果。
	Result *string `json:"result,omitempty"`

	// 工作流度量信息列表。
	MetricList *[]WorkflowMetricPair `json:"metric_list,omitempty"`
}

func (o WorkflowConditionExecution) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowConditionExecution struct{}"
	}

	return strings.Join([]string{"WorkflowConditionExecution", string(data)}, " ")
}
