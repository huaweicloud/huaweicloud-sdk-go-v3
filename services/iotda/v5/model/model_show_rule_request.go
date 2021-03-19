package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowRuleRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	RuleId string `json:"rule_id"`
}

func (o ShowRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowRuleRequest", string(data)}, " ")
}
