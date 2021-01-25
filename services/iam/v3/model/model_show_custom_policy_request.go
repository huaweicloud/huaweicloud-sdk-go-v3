package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowCustomPolicyRequest struct {
	RoleId string `json:"role_id"`
}

func (o ShowCustomPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowCustomPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowCustomPolicyRequest", string(data)}, " ")
}
