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
type CreatePolicyAssignmentsRequest struct {
	Body *PolicyAssignmentRequestBody `json:"body,omitempty"`
}

func (o CreatePolicyAssignmentsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreatePolicyAssignmentsRequest", string(data)}, " ")
}
