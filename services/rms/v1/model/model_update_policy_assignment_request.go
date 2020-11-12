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
type UpdatePolicyAssignmentRequest struct {
	PolicyAssignmentId string                       `json:"policy_assignment_id"`
	Body               *PolicyAssignmentRequestBody `json:"body,omitempty"`
}

func (o UpdatePolicyAssignmentRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UpdatePolicyAssignmentRequest", string(data)}, " ")
}
