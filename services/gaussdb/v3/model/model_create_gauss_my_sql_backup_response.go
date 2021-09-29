package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type CreateGaussMySqlBackupResponse struct {
	Backup         *Backup `json:"backup,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateGaussMySqlBackupResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlBackupResponse struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlBackupResponse", string(data)}, " ")
}
