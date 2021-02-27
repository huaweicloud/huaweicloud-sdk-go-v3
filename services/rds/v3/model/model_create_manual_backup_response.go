package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateManualBackupResponse struct {
	Backup         *BackupInfo `json:"backup,omitempty"`
	HttpStatusCode int         `json:"-"`
}

func (o CreateManualBackupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateManualBackupResponse struct{}"
	}

	return strings.Join([]string{"CreateManualBackupResponse", string(data)}, " ")
}
