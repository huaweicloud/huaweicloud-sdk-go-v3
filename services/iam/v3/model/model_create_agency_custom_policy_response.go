package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateAgencyCustomPolicyResponse struct {
	Role           *AgencyPolicyRoleResult `json:"role,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o CreateAgencyCustomPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateAgencyCustomPolicyResponse struct{}"
	}

	return strings.Join([]string{"CreateAgencyCustomPolicyResponse", string(data)}, " ")
}
