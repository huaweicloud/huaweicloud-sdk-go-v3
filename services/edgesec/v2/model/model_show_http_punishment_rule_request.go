package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowHttpPunishmentRuleRequest Request Object
type ShowHttpPunishmentRuleRequest struct {

	// 策略id
	PolicyId string `json:"policy_id"`

	// 防护规则id
	RuleId string `json:"rule_id"`
}

func (o ShowHttpPunishmentRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpPunishmentRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowHttpPunishmentRuleRequest", string(data)}, " ")
}
