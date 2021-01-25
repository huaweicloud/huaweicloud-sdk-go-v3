package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ImportBackupRequest struct {
	Body *BackupSyncReq `json:"body,omitempty"`
}

func (o ImportBackupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ImportBackupRequest struct{}"
	}

	return strings.Join([]string{"ImportBackupRequest", string(data)}, " ")
}
