package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeleteRuleActionResponse struct {
	Body           *string `json:"body,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteRuleActionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRuleActionResponse struct{}"
	}

	return strings.Join([]string{"DeleteRuleActionResponse", string(data)}, " ")
}
