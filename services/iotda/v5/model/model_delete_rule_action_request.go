package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRuleActionRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`

	ActionId string `json:"action_id"`
}

func (o DeleteRuleActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRuleActionRequest struct{}"
	}

	return strings.Join([]string{"DeleteRuleActionRequest", string(data)}, " ")
}
