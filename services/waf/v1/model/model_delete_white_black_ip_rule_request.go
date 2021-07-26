package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteWhiteBlackIpRuleRequest struct {
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// whiteblackIpRuleId

	RuleId string `json:"rule_id"`
}

func (o DeleteWhiteBlackIpRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteWhiteBlackIpRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteWhiteBlackIpRuleRequest", string(data)}, " ")
}
