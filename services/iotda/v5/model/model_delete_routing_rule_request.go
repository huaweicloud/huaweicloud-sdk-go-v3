package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRoutingRuleRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	RuleId     string  `json:"rule_id"`
}

func (o DeleteRoutingRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRoutingRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteRoutingRuleRequest", string(data)}, " ")
}
