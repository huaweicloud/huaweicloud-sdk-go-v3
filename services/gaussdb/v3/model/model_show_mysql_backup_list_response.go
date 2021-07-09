package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowMysqlBackupListResponse struct {
	// 备份信息。

	Backups *[]Backups `json:"backups,omitempty"`

	TotalCount     float32 `json:"total_count,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowMysqlBackupListResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowMysqlBackupListResponse struct{}"
	}

	return strings.Join([]string{"ShowMysqlBackupListResponse", string(data)}, " ")
}
