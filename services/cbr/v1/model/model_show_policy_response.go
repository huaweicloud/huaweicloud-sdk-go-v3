package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowPolicyResponse struct {
	Policy         *Policy `json:"policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowPolicyResponse", string(data)}, " ")
}
