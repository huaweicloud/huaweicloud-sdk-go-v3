package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MetricItem 监控指标项
type MetricItem struct {

	// **参数解释**：  监控指标键名。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Metric *string `json:"metric,omitempty"`

	// **参数解释**：  监控指标显示名称。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**：  监控指标过滤条件。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Filter *string `json:"filter,omitempty"`
}

func (o MetricItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricItem struct{}"
	}

	return strings.Join([]string{"MetricItem", string(data)}, " ")
}
