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

type RestoreTableInfo struct {
	// 旧表名
	OldName string `json:"oldName"`
	// 新表名
	NewName string `json:"newName"`
}

func (o RestoreTableInfo) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RestoreTableInfo", string(data)}, " ")
}
