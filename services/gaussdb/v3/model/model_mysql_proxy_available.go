package model

import (
	"encoding/json"

	"strings"
)

type MysqlProxyAvailable struct {
	// 可用区编码。

	Code *string `json:"code,omitempty"`
	// 可用区描述。

	Description *string `json:"description,omitempty"`
}

func (o MysqlProxyAvailable) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MysqlProxyAvailable struct{}"
	}

	return strings.Join([]string{"MysqlProxyAvailable", string(data)}, " ")
}
