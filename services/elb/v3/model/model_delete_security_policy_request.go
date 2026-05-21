package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteSecurityPolicyRequest Request Object
type DeleteSecurityPolicyRequest struct {

	// **参数解释**：自定义安全策略的ID。  **约束限制**：不涉及  **取值范围**：不涉及  **默认取值**：不涉及
	SecurityPolicyId string `json:"security_policy_id"`
}

func (o DeleteSecurityPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteSecurityPolicyRequest struct{}"
	}

	return strings.Join([]string{"DeleteSecurityPolicyRequest", string(data)}, " ")
}
