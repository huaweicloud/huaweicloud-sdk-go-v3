package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type EnablePolicyAssignmentRequest struct {
	// 规则ID

	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o EnablePolicyAssignmentRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "EnablePolicyAssignmentRequest struct{}"
	}

	return strings.Join([]string{"EnablePolicyAssignmentRequest", string(data)}, " ")
}
