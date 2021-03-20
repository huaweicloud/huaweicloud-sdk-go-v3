package model

import (
	"encoding/json"

	"strings"
)

// 协议信息
type ProtocolIdInfo struct {
	// 协议id

	Id string `json:"id"`
}

func (o ProtocolIdInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProtocolIdInfo struct{}"
	}

	return strings.Join([]string{"ProtocolIdInfo", string(data)}, " ")
}
