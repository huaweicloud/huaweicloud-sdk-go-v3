package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowRoutingRuleRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	RuleId string `json:"rule_id"`
}

func (o ShowRoutingRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRoutingRuleRequest struct{}"
	}

	return strings.Join([]string{"ShowRoutingRuleRequest", string(data)}, " ")
}
