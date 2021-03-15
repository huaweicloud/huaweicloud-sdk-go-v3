package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRoutingRuleRequest struct {
	InstanceId *string     `json:"Instance-Id,omitempty"`
	Body       *AddRuleReq `json:"body,omitempty"`
}

func (o CreateRoutingRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRoutingRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateRoutingRuleRequest", string(data)}, " ")
}
