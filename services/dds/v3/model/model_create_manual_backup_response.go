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

// Response Object
type CreateManualBackupResponse struct {
	// 手动备份ID。
	BackupId       *string `json:"backup_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateManualBackupResponse) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"CreateManualBackupResponse", string(data)}, " ")
}
