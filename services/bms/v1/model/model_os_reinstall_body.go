/*
 * BMS
 *
 * BMS Open API
 *
 */

package model

import (
	"encoding/json"

	"strings"
)

// 重装裸金属服务器操作系统接口请求结构体
type OsReinstallBody struct {
	OsReinstall *OsReinstall `json:"os-reinstall"`
}

func (o OsReinstallBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"OsReinstallBody", string(data)}, " ")
}
