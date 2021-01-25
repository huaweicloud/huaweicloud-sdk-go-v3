package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowCustomPolicyResponse struct {
	Role           *PolicyRoleResult `json:"role,omitempty"`
	HttpStatusCode int               `json:"-"`
}

func (o ShowCustomPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCustomPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowCustomPolicyResponse", string(data)}, " ")
}
