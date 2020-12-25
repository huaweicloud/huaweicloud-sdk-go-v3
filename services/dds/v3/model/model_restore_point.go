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

// 用于恢复的备份ID。
type RestorePoint struct {
	// 用于恢复的备份ID。
	BackupId string `json:"backup_id"`
}

func (o RestorePoint) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestorePoint", string(data)}, " ")
}
