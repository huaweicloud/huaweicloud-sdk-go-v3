package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoScalingPolicyResponse Response Object
type SetAutoScalingPolicyResponse struct {

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  是否开启自动变配。  **约束限制**：  不涉及。  **取值范围**：  - ON：开启自动变配 - OFF：关闭自动变配  **默认取值**：  不涉及。
	Status         *string `json:"status,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o SetAutoScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"SetAutoScalingPolicyResponse", string(data)}, " ")
}
