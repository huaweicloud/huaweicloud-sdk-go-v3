package model

import (
	"encoding/json"

	"strings"
)

type NodeItem struct {
	// 节点ID

	Uid *string `json:"uid,omitempty"`
}

func (o NodeItem) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NodeItem struct{}"
	}

	return strings.Join([]string{"NodeItem", string(data)}, " ")
}
