package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSecurityPolicyRequest Request Object
type ShowSecurityPolicyRequest struct {

	// **参数解释**：自定义安全策略的ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	SecurityPolicyId string `json:"security_policy_id"`
}

func (o ShowSecurityPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSecurityPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowSecurityPolicyRequest", string(data)}, " ")
}
