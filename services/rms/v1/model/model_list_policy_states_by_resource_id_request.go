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
type ListPolicyStatesByResourceIdRequest struct {
	ResourceId      string  `json:"resource_id"`
	ComplianceState *string `json:"compliance_state,omitempty"`
	Limit           *int32  `json:"limit,omitempty"`
	Marker          *string `json:"marker,omitempty"`
}

func (o ListPolicyStatesByResourceIdRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ListPolicyStatesByResourceIdRequest", string(data)}, " ")
}
