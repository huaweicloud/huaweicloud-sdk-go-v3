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
type DeletePolicyAssignmentResponse struct {
}

func (o DeletePolicyAssignmentResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DeletePolicyAssignmentResponse", string(data)}, " ")
}
