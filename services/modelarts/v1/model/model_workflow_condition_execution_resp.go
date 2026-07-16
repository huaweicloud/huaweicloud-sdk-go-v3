package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkflowConditionExecutionResp workflow condition execution
type WorkflowConditionExecutionResp struct {

	// **参数解释**：执行结果。 **取值范围**：不涉及。
	Result *string `json:"result,omitempty"`

	// **参数解释**：工作流度量信息列表。
	MetricList *[]WorkflowMetricPairResp `json:"metric_list,omitempty"`
}

func (o WorkflowConditionExecutionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkflowConditionExecutionResp struct{}"
	}

	return strings.Join([]string{"WorkflowConditionExecutionResp", string(data)}, " ")
}
