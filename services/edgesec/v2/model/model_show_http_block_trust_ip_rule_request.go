package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpBlockTrustIpRuleRequest Request Object
type ShowHttpBlockTrustIpRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`
}

func (o ShowHttpBlockTrustIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpBlockTrustIpRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpBlockTrustIpRuleRequest", string(data)}, " ")
}
