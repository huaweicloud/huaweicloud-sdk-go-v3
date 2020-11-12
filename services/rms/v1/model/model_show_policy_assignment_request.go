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
type ShowPolicyAssignmentRequest struct {
	PolicyAssignmentId string `json:"policy_assignment_id"`
}

func (o ShowPolicyAssignmentRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowPolicyAssignmentRequest", string(data)}, " ")
}
