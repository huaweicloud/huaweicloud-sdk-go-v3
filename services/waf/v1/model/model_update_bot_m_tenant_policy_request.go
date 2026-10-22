package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBotMTenantPolicyRequest Request Object
type UpdateBotMTenantPolicyRequest struct {

	// policyid
	PolicyId string `json:"policy_id"`

	Body *UpdateBotMTenantPolicyRequestBody `json:"body,omitempty"`
}

func (o UpdateBotMTenantPolicyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBotMTenantPolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdateBotMTenantPolicyRequest", string(data)}, " ")
}
