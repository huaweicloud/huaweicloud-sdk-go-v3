package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateRuleRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	RuleId string `json:"rule_id"`

	Body *Rule `json:"body,omitempty"`
}

func (o UpdateRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateRuleRequest", string(data)}, " ")
}
