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
type ListPolicyAssignmentsRequest struct {
}

func (o ListPolicyAssignmentsRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPolicyAssignmentsRequest", string(data)}, " ")
}
