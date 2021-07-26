package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateWhiteblackipRuleRequest struct {
	// 策略id（策略id从查询防护策略列表接口获取）

	PolicyId string `json:"policy_id"`

	Body *CreateWhiteBlackIpRuleRequestBody `json:"body,omitempty"`
}

func (o CreateWhiteblackipRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateWhiteblackipRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateWhiteblackipRuleRequest", string(data)}, " ")
}
