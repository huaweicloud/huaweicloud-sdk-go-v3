package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListWorkflowStepExecutionResponse Response Object
type ListWorkflowStepExecutionResponse struct {

	// **参数解释**：总数。 **取值范围**：不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：返回个数。 **取值范围**：不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**：StepExecution数组。
	Items *[]StepExecutionResp `json:"items,omitempty"`

	// **参数解释**：默认排序。 **取值范围**：不涉及。
	DefaultOrder *string `json:"default_order,omitempty"`

	CompareColumns *CompareColumns `json:"compare_columns,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListWorkflowStepExecutionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWorkflowStepExecutionResponse struct{}"
	}

	return strings.Join([]string{"ListWorkflowStepExecutionResponse", string(data)}, " ")
}
