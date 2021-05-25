package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowPolicyAssignmentRequest struct {
	// 规则ID

	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o ShowPolicyAssignmentRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowPolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"ShowPolicyAssignmentRequest", string(data)}, " ")
}
