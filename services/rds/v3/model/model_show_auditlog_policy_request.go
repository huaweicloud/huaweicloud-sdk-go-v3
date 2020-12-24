/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowAuditlogPolicyRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	InstanceId string  `json:"instance_id"`
}

func (o ShowAuditlogPolicyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowAuditlogPolicyRequest", string(data)}, " ")
}
