package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdatePolicyRuleStatusResponse struct {
	HttpStatusCode int `json:"-"`
}

func (o UpdatePolicyRuleStatusResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePolicyRuleStatusResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRuleStatusResponse", string(data)}, " ")
}
