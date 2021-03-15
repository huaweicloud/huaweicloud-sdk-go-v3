package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateRoutingRuleRequest struct {
	InstanceId *string        `json:"Instance-Id,omitempty"`
	RuleId     string         `json:"rule_id"`
	Body       *UpdateRuleReq `json:"body,omitempty"`
}

func (o UpdateRoutingRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRoutingRuleRequest struct{}"
	}

	return strings.Join([]string{"UpdateRoutingRuleRequest", string(data)}, " ")
}
