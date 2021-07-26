package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteAntitamperRuleRequest struct {
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// antitamperRuleId

	RuleId string `json:"rule_id"`
}

func (o DeleteAntitamperRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteAntitamperRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteAntitamperRuleRequest", string(data)}, " ")
}
