package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowServerlessScalingPolicyResponse Response Object
type ShowServerlessScalingPolicyResponse struct {

	// **参数描述**：  自定义扩容步长。  **约束限制**：  不涉及。  **取值范围**：  2-算力上限的一半。  **默认取值**：  不涉及。
	EnlargeStepSize *string `json:"enlarge_step_size,omitempty"`

	CustomScalingConfig *CustomScalingConfig `json:"custom_scaling_config,omitempty"`
	HttpStatusCode      int                  `json:"-"`
}

func (o ShowServerlessScalingPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowServerlessScalingPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowServerlessScalingPolicyResponse", string(data)}, " ")
}
