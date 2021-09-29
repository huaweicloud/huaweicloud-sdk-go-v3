package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateGaussMySqlBackupRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`

	Body *MysqlCreateBackupRequest `json:"body,omitempty"`
}

func (o CreateGaussMySqlBackupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateGaussMySqlBackupRequest struct{}"
	}

	return strings.Join([]string{"CreateGaussMySqlBackupRequest", string(data)}, " ")
}
