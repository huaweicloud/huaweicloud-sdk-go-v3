package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMysqlBackupPolicyResponse struct {
	BackupPolicy   *BackupPolicy `json:"backup_policy,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowMysqlBackupPolicyResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlBackupPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowMysqlBackupPolicyResponse", string(data)}, " ")
}
