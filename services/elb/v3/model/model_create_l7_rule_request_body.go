package model

import (
	"encoding/json"

	"strings"
)

// This is a auto create Body Object
type CreateL7RuleRequestBody struct {
	Rule *CreateRuleOption `json:"rule"`
}

func (o CreateL7RuleRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateL7RuleRequestBody struct{}"
	}

	return strings.Join([]string{"CreateL7RuleRequestBody", string(data)}, " ")
}
