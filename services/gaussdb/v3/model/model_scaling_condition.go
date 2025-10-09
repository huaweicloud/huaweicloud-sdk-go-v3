package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ScalingCondition **参数描述**：  变配条件。  **约束限制**：  至少各定义一个CPU使用率和内存使用率指标条件，最多各定义两个。
type ScalingCondition struct {

	// **参数描述**:  变配条件名称。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字和中划线组成，且长度不超过32个字符，不能为空。  **默认取值**：  不涉及。
	ConditionId *string `json:"condition_id,omitempty"`

	// **参数描述**：  持续时间，单位是毫秒。  **约束限制**：  不涉及。  **取值范围**：  2000-5000。  **默认取值**：  不涉及。
	Duration *int32 `json:"duration,omitempty"`

	// **参数解释**:  间隔时间，单位是毫秒。  **约束限制**：  不涉及。  **取值范围**：  1000-5000。  **默认取值**：  不涉及。
	Interval *int32 `json:"interval,omitempty"`

	// **参数描述**:  指标变配条件。  **约束条件**：  不涉及。
	MetricConditions *[]MetricCondition `json:"metric_conditions,omitempty"`
}

func (o ScalingCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ScalingCondition struct{}"
	}

	return strings.Join([]string{"ScalingCondition", string(data)}, " ")
}
