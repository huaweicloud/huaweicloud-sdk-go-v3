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
type ShowBackupPolicyRequest struct {
	XLanguage  *string `json:"X-Language,omitempty"`
	InstanceId string  `json:"instance_id"`
}

func (o ShowBackupPolicyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"ShowBackupPolicyRequest", string(data)}, " ")
}
