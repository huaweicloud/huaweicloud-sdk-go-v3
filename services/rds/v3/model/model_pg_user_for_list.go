package model

import (
	"encoding/json"

	"strings"
)

// 数据库用户信息。
type PgUserForList struct {
	// 数据库用户名称。
	Name string `json:"name"`
	// 用户的权限属性。
	Attributes *interface{} `json:"attributes,omitempty"`
	// 用户的权限属性。
	Memberof *[]string `json:"memberof,omitempty"`
}

func (o PgUserForList) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PgUserForList struct{}"
	}

	return strings.Join([]string{"PgUserForList", string(data)}, " ")
}
