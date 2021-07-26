package model

import (
	"encoding/json"

	"strings"
)

// 类型
type RouteServerBody struct {
	// 协议

	BackProtocol *string `json:"back_protocol,omitempty"`
	// ip信息

	Address *string `json:"address,omitempty"`
	// 端口信息

	Port *int32 `json:"port,omitempty"`
}

func (o RouteServerBody) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RouteServerBody struct{}"
	}

	return strings.Join([]string{"RouteServerBody", string(data)}, " ")
}
