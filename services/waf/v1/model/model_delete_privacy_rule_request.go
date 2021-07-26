package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePrivacyRuleRequest struct {
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// privacyRuleId

	RuleId string `json:"rule_id"`
}

func (o DeletePrivacyRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePrivacyRuleRequest struct{}"
	}

	return strings.Join([]string{"DeletePrivacyRuleRequest", string(data)}, " ")
}
