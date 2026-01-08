package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateLoadbalancerAutoscalingOption struct {

	// **参数解释**：负载均衡器实例弹性扩缩容开关。  **约束限制**：网关型LB不支持该字段。  **取值范围**： - true：开启。 - false：关闭。   **默认取值**：不涉及
	Enable bool `json:"enable"`

	// **参数解释**：实例弹性扩缩容的最小七层规格ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	MinL7FlavorId *string `json:"min_l7_flavor_id,omitempty"`
}

func (o CreateLoadbalancerAutoscalingOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadbalancerAutoscalingOption struct{}"
	}

	return strings.Join([]string{"CreateLoadbalancerAutoscalingOption", string(data)}, " ")
}
