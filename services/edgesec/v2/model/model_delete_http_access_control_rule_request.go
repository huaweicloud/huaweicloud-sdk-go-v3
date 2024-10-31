package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpAccessControlRuleRequest Request Object
type DeleteHttpAccessControlRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`
}

func (o DeleteHttpAccessControlRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpAccessControlRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteHttpAccessControlRuleRequest", string(data)}, " ")
}
