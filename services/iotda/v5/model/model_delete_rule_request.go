package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRuleRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	RuleId     string  `json:"rule_id"`
}

func (o DeleteRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteRuleRequest", string(data)}, " ")
}
