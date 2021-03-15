package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowRuleActionRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	ActionId   string  `json:"action_id"`
}

func (o ShowRuleActionRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowRuleActionRequest struct{}"
	}

	return strings.Join([]string{"ShowRuleActionRequest", string(data)}, " ")
}
