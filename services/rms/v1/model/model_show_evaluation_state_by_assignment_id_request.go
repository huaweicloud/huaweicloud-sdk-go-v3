/*
 * RMS
 *
 * Resource Manager Api
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowEvaluationStateByAssignmentIdRequest struct {
	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o ShowEvaluationStateByAssignmentIdRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowEvaluationStateByAssignmentIdRequest", string(data)}, " ")
}
