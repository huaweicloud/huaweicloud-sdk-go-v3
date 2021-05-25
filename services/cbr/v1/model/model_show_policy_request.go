package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPolicyRequest struct {
	// 策略ID

	PolicyId string `json:"policy_id"`
}

func (o ShowPolicyRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPolicyRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyRequest", string(data)}, " ")
}
