package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteBackupRequest struct {
	BackupId string `json:"backup_id"`
}

func (o DeleteBackupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBackupRequest struct{}"
	}

	return strings.Join([]string{"DeleteBackupRequest", string(data)}, " ")
}
