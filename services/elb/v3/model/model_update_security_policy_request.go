package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateSecurityPolicyRequest Request Object
type UpdateSecurityPolicyRequest struct {

	// 自定义安全策略的ID。
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
