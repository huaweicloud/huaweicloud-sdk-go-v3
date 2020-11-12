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

// Response Object
type RunEvaluationByPolicyAssignmentIdResponse struct {
}

func (o RunEvaluationByPolicyAssignmentIdResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RunEvaluationByPolicyAssignmentIdResponse", string(data)}, " ")
}
