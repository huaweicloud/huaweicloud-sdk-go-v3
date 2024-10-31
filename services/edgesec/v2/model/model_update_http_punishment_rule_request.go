package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHttpPunishmentRuleRequest Request Object
type UpdateHttpPunishmentRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`

	Body *UpdateHttpPunishmentRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateHttpPunishmentRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHttpPunishmentRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateHttpPunishmentRuleRequest", string(data)}, " ")
}
