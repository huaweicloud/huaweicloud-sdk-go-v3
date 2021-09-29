package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowGaussMySqlBackupListResponse struct {
	// 备份信息。

	Backups *[]Backups `json:"backups,omitempty"`
	// 备份文件的总数。

	TotalCount     *int64 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowGaussMySqlBackupListResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowGaussMySqlBackupListResponse struct{}"
	}

	return strings.Join([]string{"ShowGaussMySqlBackupListResponse", string(data)}, " ")
}
