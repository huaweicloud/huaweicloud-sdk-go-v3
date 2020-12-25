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

type SetBackupPolicyRequestBody struct {
	BackupPolicy *BackupPolicy `json:"backup_policy"`
}

func (o SetBackupPolicyRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"SetBackupPolicyRequestBody", string(data)}, " ")
}
