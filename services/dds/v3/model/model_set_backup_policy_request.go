/*
 * DDS
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
type SetBackupPolicyRequest struct {
	InstanceId string                      `json:"instance_id"`
	Body       *SetBackupPolicyRequestBody `json:"body,omitempty"`
}

func (o SetBackupPolicyRequest) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SetBackupPolicyRequest", string(data)}, " ")
}
