package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReadOnlyScalingStrategy 只读自动变配策略。
type ReadOnlyScalingStrategy struct {

	// **参数解释**：  只读扩容开关。  **约束限制**：  不涉及。  **取值范围**：  - ON：开启 - OFF：关闭  **默认取值**：  不涉及。
	ReadOnlyEnlargeEnabled *string `json:"read_only_enlarge_enabled,omitempty"`

	// **参数解释**：  只读缩容开关。  **约束限制**：  不涉及。  **取值范围**：  - ON：开启 - OFF：关闭  **默认取值**：  不涉及。
	ReadOnlyReduceEnabled *string `json:"read_only_reduce_enabled,omitempty"`

	// **参数解释**：  观测窗口时间，单位秒。  **约束限制**：  不涉及。  **取值范围**：  - 120 - 300 - 600 - 900 - 1800  **默认取值**：  不涉及。
	ReadOnlyMonitorCycle *string `json:"read_only_monitor_cycle,omitempty"`

	// **参数解释**：  静默期，单位秒。  **约束限制**：  不涉及。  **取值范围**：  - 300 - 600 - 1800 - 3600 - 7200 - 10800 - 86400 - 604800  **默认取值**：  不涉及。
	ReadOnlySilenceCycle *string `json:"read_only_silence_cycle,omitempty"`

	// **参数解释**：  只读最大节点数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	MaxReadOnlyCount *string `json:"max_read_only_count,omitempty"`

	// **参数解释**：  只读扩容阈值。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ReadOnlyEnlargeThreshold *string `json:"read_only_enlarge_threshold,omitempty"`

	// **参数解释**：  扩容新增只读规格。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ReadOnlyFlavor *string `json:"read_only_flavor,omitempty"`

	// **参数解释**：  只读最小节点数。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	MinReadOnlyCount *string `json:"min_read_only_count,omitempty"`

	// **参数解释**：  只读缩容阈值。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ReadOnlyReduceThreshold *string `json:"read_only_reduce_threshold,omitempty"`
}

func (o ReadOnlyScalingStrategy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReadOnlyScalingStrategy struct{}"
	}

	return strings.Join([]string{"ReadOnlyScalingStrategy", string(data)}, " ")
}
