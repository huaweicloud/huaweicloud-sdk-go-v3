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
type DisablePolicyAssignmentResponse struct {
}

func (o DisablePolicyAssignmentResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"DisablePolicyAssignmentResponse", string(data)}, " ")
}
