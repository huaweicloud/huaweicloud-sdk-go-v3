package model

import (
	"encoding/json"

	"strings"
)

type Node struct {
	// 节点EIP信息
	IpPort *string `json:"ip_port,omitempty"`
	// 节点所在通道数组
	Channels *[]string `json:"channels,omitempty"`
}

func (o Node) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "Node struct{}"
	}

	return strings.Join([]string{"Node", string(data)}, " ")
}
