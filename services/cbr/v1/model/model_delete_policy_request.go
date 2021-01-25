package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeletePolicyRequest struct {
	PolicyId string `json:"policy_id"`
}

func (o DeletePolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeletePolicyRequest struct{}"
	}

	return strings.Join([]string{"DeletePolicyRequest", string(data)}, " ")
}
