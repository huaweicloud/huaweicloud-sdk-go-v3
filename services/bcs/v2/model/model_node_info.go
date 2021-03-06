package model

import (
	"encoding/json"

	"strings"
)

// 节点信息
type NodeInfo struct {
	// 节点名

	Name *string `json:"name,omitempty"`
	// 节点ID

	Id *string `json:"id,omitempty"`
}

func (o NodeInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NodeInfo struct{}"
	}

	return strings.Join([]string{"NodeInfo", string(data)}, " ")
}
