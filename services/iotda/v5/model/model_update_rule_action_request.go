package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateRuleActionRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	ActionId string `json:"action_id"`

	Body *UpdateActionReq `json:"body,omitempty"`
}

func (o UpdateRuleActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateRuleActionRequest struct{}"
	}

	return strings.Join([]string{"UpdateRuleActionRequest", string(data)}, " ")
}
