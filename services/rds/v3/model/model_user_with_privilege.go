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

// 用户及其权限。
type UserWithPrivilege struct {
	// 用户名。
	Name string `json:"name"`
	// 是否为只读权限。
	Readonly bool `json:"readonly"`
}

func (o UserWithPrivilege) String() string {
	data, _ := json.Marshal(o)
	return strings.Join([]string{"UserWithPrivilege", string(data)}, " ")
}
