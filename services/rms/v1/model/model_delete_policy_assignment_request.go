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
type DeletePolicyAssignmentRequest struct {
	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o DeletePolicyAssignmentRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePolicyAssignmentRequest", string(data)}, " ")
}
