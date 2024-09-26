package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpPolicyRuleStatusRequest Request Object
type UpdateHttpPolicyRuleStatusRequest struct {

	// 防护策略id
	PolicyId string `json:"policy_id"`

	// 防护策略规则 cc|custom|whiteblackip|privacy|ignore|geoip|antitamper|antileakage|punishment|ip-reputation
	RuleType string `json:"rule_type"`

	// 防护策略规则id
	RuleId string `json:"rule_id"`

	Body *UpdateHttpPolicyRuleStatusRequestBody `json:"body,omitempty"`
}

func (o UpdateHttpPolicyRuleStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpPolicyRuleStatusRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpPolicyRuleStatusRequest", string(data)}, " ")
}
