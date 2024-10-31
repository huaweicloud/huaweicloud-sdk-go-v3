package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpBlockTrustIpRuleRequest Request Object
type UpdateHttpBlockTrustIpRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`

	Body *UpdateHttpBlockTrustIpRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateHttpBlockTrustIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpBlockTrustIpRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpBlockTrustIpRuleRequest", string(data)}, " ")
}
