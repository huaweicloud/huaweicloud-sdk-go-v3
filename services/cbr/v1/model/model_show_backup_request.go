package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBackupRequest struct {
	// 备份ID

	BackupId string `json:"backup_id"`
}

func (o ShowBackupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBackupRequest struct{}"
	}

	return strings.Join([]string{"ShowBackupRequest", string(data)}, " ")
}
