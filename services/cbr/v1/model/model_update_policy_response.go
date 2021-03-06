package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdatePolicyResponse struct {
	Policy         *Policy `json:"policy,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdatePolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdatePolicyResponse", string(data)}, " ")
}
