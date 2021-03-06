package model

import (
	"encoding/json"

	"strings"
)

//
type ServerRemoteConsole struct {
	// 远程登录的协议。

	Protocol string `json:"protocol"`
	// 远程登录的类型。

	Type string `json:"type"`
	// 远程登录的url。

	Url string `json:"url"`
}

func (o ServerRemoteConsole) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ServerRemoteConsole struct{}"
	}

	return strings.Join([]string{"ServerRemoteConsole", string(data)}, " ")
}
