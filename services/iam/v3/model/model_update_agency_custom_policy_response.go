package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type UpdateAgencyCustomPolicyResponse struct {
	Role           *AgencyPolicyRoleResult `json:"role,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o UpdateAgencyCustomPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAgencyCustomPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateAgencyCustomPolicyResponse", string(data)}, " ")
}
