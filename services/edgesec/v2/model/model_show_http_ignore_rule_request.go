package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpIgnoreRuleRequest Request Object
type ShowHttpIgnoreRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`
}

func (o ShowHttpIgnoreRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpIgnoreRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpIgnoreRuleRequest", string(data)}, " ")
}
