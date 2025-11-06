package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAutoScalingPolicyResponse Response Object
type ListAutoScalingPolicyResponse struct {

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  是否开启autoScaling，OFF为关闭，ON为开启  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**：  观察窗口，单位秒  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	MonitorCycle *int32 `json:"monitor_cycle,omitempty"`

	// **参数解释**：  静默期，单位秒  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	SilenceCycle *int32 `json:"silence_cycle,omitempty"`

	// **参数解释**：  自动升配触发阈值，单位百分比  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	EnlargeThreshold *int32 `json:"enlarge_threshold,omitempty"`

	// **参数解释**：  最大变配规格上限  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	MaxFlavor *string `json:"max_flavor,omitempty"`

	// **参数解释**：  自动降配状态，on为自动降配开启；off为关闭  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ReduceEnabled *string `json:"reduce_enabled,omitempty"`

	// **参数解释**：  自动降配触发阈值  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	ReduceThreshold *int32 `json:"reduce_threshold,omitempty"`

	// **参数解释**：  最大变配规格下限  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	MinFlavor      *string `json:"min_flavor,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListAutoScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAutoScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ListAutoScalingPolicyResponse", string(data)}, " ")
}
