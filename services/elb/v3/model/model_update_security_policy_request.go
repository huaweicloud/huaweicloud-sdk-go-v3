package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityPolicyRequest Request Object
type UpdateSecurityPolicyRequest struct {

	// **参数解释**：自定义安全策略的ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	SecurityPolicyId string `json:"security_policy_id"`

	Body *UpdateSecurityPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateSecurityPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSecurityPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateSecurityPolicyRequest", string(data)}, " ")
}
