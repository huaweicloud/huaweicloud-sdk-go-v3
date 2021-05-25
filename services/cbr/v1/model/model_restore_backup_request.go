package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type RestoreBackupRequest struct {
	// 备份id

	BackupId string `json:"backup_id"`

	Body *BackupRestoreReq `json:"body,omitempty"`
}

func (o RestoreBackupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RestoreBackupRequest struct{}"
	}

	return strings.Join([]string{"RestoreBackupRequest", string(data)}, " ")
}
