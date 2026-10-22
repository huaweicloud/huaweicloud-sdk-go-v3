package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBotMRuleRequest Request Object
type ShowBotMRuleRequest struct {

	// policyid
	PolicyId string `json:"policy_id"`

	// ruleid
	RuleId string `json:"rule_id"`
}

func (o ShowBotMRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBotMRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowBotMRuleRequest", string(data)}, " ")
}
