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
type DisablePolicyAssignmentRequest struct {
	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o DisablePolicyAssignmentRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DisablePolicyAssignmentRequest", string(data)}, " ")
}
