package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowEvaluationStateByAssignmentIdRequest struct {
	// 规则ID

	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o ShowEvaluationStateByAssignmentIdRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowEvaluationStateByAssignmentIdRequest struct{}"
	}

	return strings.Join([]string{"ShowEvaluationStateByAssignmentIdRequest", string(data)}, " ")
}
