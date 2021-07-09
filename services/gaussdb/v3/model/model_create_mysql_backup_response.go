package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateMysqlBackupResponse struct {
	Backup         *Backup `json:"backup,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateMysqlBackupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMysqlBackupResponse struct{}"
	}

	return strings.Join([]string{"CreateMysqlBackupResponse", string(data)}, " ")
}
