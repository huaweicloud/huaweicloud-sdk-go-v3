package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteL7ruleRequest struct {
	// 转发策略id

	L7policyId string `json:"l7policy_id"`
	// 转发规则id

	L7ruleId string `json:"l7rule_id"`
}

func (o DeleteL7ruleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteL7ruleRequest struct{}"
	}

	return strings.Join([]string{"DeleteL7ruleRequest", string(data)}, " ")
}
