package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeRuleStatusRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	RuleId string `json:"rule_id"`

	Body *RuleStatus `json:"body,omitempty"`
}

func (o ChangeRuleStatusRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeRuleStatusRequest struct{}"
	}

	return strings.Join([]string{"ChangeRuleStatusRequest", string(data)}, " ")
}
