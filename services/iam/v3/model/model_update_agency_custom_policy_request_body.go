package model

import (
	"encoding/json"

	"strings"
)

//
type UpdateAgencyCustomPolicyRequestBody struct {
	Role *AgencyPolicyRoleOption `json:"role"`
}

func (o UpdateAgencyCustomPolicyRequestBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateAgencyCustomPolicyRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateAgencyCustomPolicyRequestBody", string(data)}, " ")
}
