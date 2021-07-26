package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type DeletePrivacyRuleResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o DeletePrivacyRuleResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePrivacyRuleResponse struct{}"
	}

	return strings.Join([]string{"DeletePrivacyRuleResponse", string(data)}, " ")
}
