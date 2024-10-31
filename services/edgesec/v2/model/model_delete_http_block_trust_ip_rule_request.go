package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpBlockTrustIpRuleRequest Request Object
type DeleteHttpBlockTrustIpRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`
}

func (o DeleteHttpBlockTrustIpRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpBlockTrustIpRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteHttpBlockTrustIpRuleRequest", string(data)}, " ")
}
