package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRuleActionRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	XLBService *string `json:"x-LB-Service,omitempty"`

	Body *AddActionReq `json:"body,omitempty"`
}

func (o CreateRuleActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRuleActionRequest struct{}"
	}

	return strings.Join([]string{"CreateRuleActionRequest", string(data)}, " ")
}
