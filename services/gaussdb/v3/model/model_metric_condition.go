package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetricCondition struct {

	// **参数描述**:  指标名称。  **约束限制**：  不涉及。  **取值范围**： - cpuTotalUsage：CPU使用率。 - memoryTotalUsage：内存使用率。  **默认取值**：  不涉及。
	MetricName *string `json:"metric_name,omitempty"`

	// **参数描述**：  指标变配阈值。  **约束限制**：  取值为百分比的10000倍，比如50%对应的参数值为5000。  **取值范围**：  6000-8000。  **默认取值**：  不涉及。
	MetricValue *int32 `json:"metric_value,omitempty"`

	// **参数描述**：  比较模式。  **约束限制**：  不涉及。。  **取值范围**：  GT：大于。  **默认取值**：  不涉及。
	CompareMode *string `json:"compare_mode,omitempty"`
}

func (o MetricCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetricCondition struct{}"
	}

	return strings.Join([]string{"MetricCondition", string(data)}, " ")
}
