package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DisablePolicyAssignmentRequest struct {
	// 规则ID

	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o DisablePolicyAssignmentRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DisablePolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"DisablePolicyAssignmentRequest", string(data)}, " ")
}
