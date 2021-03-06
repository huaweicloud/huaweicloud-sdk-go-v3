package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdatePolicyRequest struct {
	// 策略ID

	PolicyId string `json:"policy_id"`

	Body *PolicyUpdateReq `json:"body,omitempty"`
}

func (o UpdatePolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdatePolicyRequest struct{}"
	}

	return strings.Join([]string{"UpdatePolicyRequest", string(data)}, " ")
}
