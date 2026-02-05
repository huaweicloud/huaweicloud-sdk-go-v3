package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSlowLogAnalysisRequest Request Object
type ListSlowLogAnalysisRequest struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in09，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  开始时间。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	StartDate float32 `json:"start_date"`

	// **参数解释**：  结束时间。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EndDate float32 `json:"end_date"`

	// **参数解释**：  分页参数: 起始值。  **约束限制**：  不涉及。  **取值范围**：  大于等于0。  **默认取值**：  默认值是0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  分页参数: 每页记录数。  **约束限制**：  不涉及。  **取值范围**：  大于0且小于等于128。  **默认取值**：  默认值是10。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：  排序。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Order *string `json:"order,omitempty"`
}

func (o ListSlowLogAnalysisRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSlowLogAnalysisRequest struct{}"
	}

	return strings.Join([]string{"ListSlowLogAnalysisRequest", string(data)}, " ")
}
