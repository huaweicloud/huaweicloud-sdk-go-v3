package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CreateMysqlBackupRequest struct {
	// 语言

	XLanguage *string `json:"X-Language,omitempty"`

	Body *MysqlCreateBackupRequest `json:"body,omitempty"`
}

func (o CreateMysqlBackupRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateMysqlBackupRequest struct{}"
	}

	return strings.Join([]string{"CreateMysqlBackupRequest", string(data)}, " ")
}
