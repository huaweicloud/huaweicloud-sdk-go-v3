package model

import (
	"encoding/json"

	"strings"
)

type RemoveNodesSpec struct {
	Login *Login `json:"login"`
	// 待操作节点列表

	Nodes *[]NodeItem `json:"nodes,omitempty"`
}

func (o RemoveNodesSpec) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "RemoveNodesSpec struct{}"
	}

	return strings.Join([]string{"RemoveNodesSpec", string(data)}, " ")
}
