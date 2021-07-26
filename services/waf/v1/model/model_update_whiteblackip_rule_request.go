package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateWhiteblackipRuleRequest struct {
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`
	// whiteblackIpRuleId

	RuleId string `json:"rule_id"`

	Body *UpdateWhiteBlackIpRuleRequestBody `json:"body,omitempty"`
}

func (o UpdateWhiteblackipRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateWhiteblackipRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateWhiteblackipRuleRequest", string(data)}, " ")
}
