package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBotMRuleActionRequest Request Object
type UpdateBotMRuleActionRequest struct {

	// policyId
	PolicyId string `json:"policy_id"`

	// 规则Id
	RuleId string `json:"rule_id"`

	Body *UpdateBotMRuleActionRequestBody `json:"body,omitempty"`
}

func (o UpdateBotMRuleActionRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBotMRuleActionRequest struct{}"
	}

	return strings.Join([]string{"UpdateBotMRuleActionRequest", string(data)}, " ")
}
