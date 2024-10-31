package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpAccessControlRuleRequest Request Object
type ShowHttpAccessControlRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`
}

func (o ShowHttpAccessControlRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpAccessControlRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpAccessControlRuleRequest", string(data)}, " ")
}
