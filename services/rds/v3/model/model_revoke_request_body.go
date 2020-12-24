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

type RevokeRequestBody struct {
	// 数据库名称。
	DbName string `json:"db_name"`
	// 解除授权的用户列表。
	Users []RevokeRequestBodyUsers `json:"users"`
}

func (o RevokeRequestBody) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"RevokeRequestBody", string(data)}, " ")
}
