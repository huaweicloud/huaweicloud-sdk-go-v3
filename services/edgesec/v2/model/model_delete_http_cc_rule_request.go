package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteHttpCcRuleRequest Request Object
type DeleteHttpCcRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// cc规则id
	RuleId string `json:"rule_id"`
}

func (o DeleteHttpCcRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteHttpCcRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteHttpCcRuleRequest", string(data)}, " ")
}
