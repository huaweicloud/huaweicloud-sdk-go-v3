package model

import (
	"encoding/json"

	"strings"
)

//
type Domains struct {
	// 是否启用账号，true为启用，false为停用，默认为true。

	Enabled bool `json:"enabled"`
	// 账号ID。

	Id string `json:"id"`
	// 账号名。

	Name string `json:"name"`

	Links *LinksSelf `json:"links"`
	// 账号的描述信息。

	Description string `json:"description"`
}

func (o Domains) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Domains struct{}"
	}

	return strings.Join([]string{"Domains", string(data)}, " ")
}
