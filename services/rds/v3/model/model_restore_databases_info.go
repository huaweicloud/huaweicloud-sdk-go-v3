/*
 * RDS
 *
 * API v3
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

type RestoreDatabasesInfo struct {
	// 库名
	Database string `json:"database"`
	// 表信息
	Tables []RestoreTableInfo `json:"tables"`
}

func (o RestoreDatabasesInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestoreDatabasesInfo", string(data)}, " ")
}
