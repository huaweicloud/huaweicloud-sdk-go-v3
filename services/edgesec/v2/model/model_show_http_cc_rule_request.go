package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpCcRuleRequest Request Object
type ShowHttpCcRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// cc规则id
	RuleId string `json:"rule_id"`
}

func (o ShowHttpCcRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpCcRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpCcRuleRequest", string(data)}, " ")
}
