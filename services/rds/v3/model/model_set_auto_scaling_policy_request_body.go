package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoScalingPolicyRequestBody 设置自动变配策略请求体。
type SetAutoScalingPolicyRequestBody struct {

	// **参数解释**：  是否开启自动变配。  **约束限制**：  不涉及。  **取值范围**：  - ON：开启自动变配 - OFF：关闭自动变配  **默认取值**：  不涉及。
	Status string `json:"status"`

	// **参数解释**：  观察窗口，单位秒。  **约束限制**：  不涉及。  **取值范围**：  300-1800  **默认取值**：  不涉及。
	MonitorCycle *int32 `json:"monitor_cycle,omitempty"`

	// **参数解释**：  静默期，单位秒。  **约束限制**：  不涉及。  **取值范围**：  300-604800  **默认取值**：  不涉及。
	SilenceCycle *int32 `json:"silence_cycle,omitempty"`

	// **参数解释**：  自动升配触发阈值，单位百分比。  **约束限制**：  不涉及。  **取值范围**：  50-100  **默认取值**：  不涉及。
	EnlargeThreshold *int32 `json:"enlarge_threshold,omitempty"`

	// **参数解释**：  最大变配规格上限。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	MaxFlavor *string `json:"max_flavor,omitempty"`

	// **参数解释**：  自动降配状态。  **约束限制**：  不涉及。  **取值范围**：  - ON：自动降配开启 - OFF：自动降配关闭  **默认取值**：  不涉及。
	ReduceEnabled *string `json:"reduce_enabled,omitempty"`

	// **参数解释**：  自动降配触发阈值。  **约束限制**：  不涉及。  **取值范围**：  10-30  **默认取值**：  不涉及。
	ReduceThreshold *int32 `json:"reduce_threshold,omitempty"`

	// **参数解释**：  最小变配规格下限。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	MinFlavor *string `json:"min_flavor,omitempty"`

	ReadOnlyScalingStrategy *ReadOnlyScalingStrategy `json:"read_only_scaling_strategy,omitempty"`
}

func (o SetAutoScalingPolicyRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoScalingPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"SetAutoScalingPolicyRequestBody", string(data)}, " ")
}
