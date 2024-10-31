package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpIgnoreRuleRequest Request Object
type DeleteHttpIgnoreRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`
}

func (o DeleteHttpIgnoreRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpIgnoreRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteHttpIgnoreRuleRequest", string(data)}, " ")
}
