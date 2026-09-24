package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SetAutoScalingPolicyRequest Request Object
type SetAutoScalingPolicyRequest struct {

	// **参数解释**：  实例ID。  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  请求语言类型。  **约束限制**：  不涉及。  **取值范围**：  - en-us - zh-cn  **默认取值**：  en-us。
	XLanguage *string `json:"X-Language,omitempty"`

	Body *SetAutoScalingPolicyRequestBody `json:"body,omitempty"`
}

func (o SetAutoScalingPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SetAutoScalingPolicyRequest struct{}"
	}

	return strings.Join([]string{"SetAutoScalingPolicyRequest", string(data)}, " ")
}
