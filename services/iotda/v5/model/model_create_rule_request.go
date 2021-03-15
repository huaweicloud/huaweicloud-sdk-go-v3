package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateRuleRequest struct {
	InstanceId *string `json:"Instance-Id,omitempty"`
	Body       *Rule   `json:"body,omitempty"`
}

func (o CreateRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRuleRequest struct{}"
	}

	return strings.Join([]string{"CreateRuleRequest", string(data)}, " ")
}
