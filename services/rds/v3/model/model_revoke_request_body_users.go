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

type RevokeRequestBodyUsers struct {
	// 数据库用户名称。
	Name string `json:"name"`
}

func (o RevokeRequestBodyUsers) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RevokeRequestBodyUsers", string(data)}, " ")
}
