package model

import (
	"encoding/json"

	"strings"
)

type PrivateNameServer struct {
	// 优先级。如果priority的值为“1”，表示会第一个采用该域名服务器进行解析。
	Priority *int32 `json:"priority,omitempty"`
	// DNS服务器地址。
	Address *string `json:"address,omitempty"`
}

func (o PrivateNameServer) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "PrivateNameServer struct{}"
	}

	return strings.Join([]string{"PrivateNameServer", string(data)}, " ")
}
