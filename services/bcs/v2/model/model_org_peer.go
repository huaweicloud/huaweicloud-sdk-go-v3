package model

import (
	"encoding/json"

	"strings"
)

type OrgPeer struct {
	// 组织名称

	Name string `json:"name"`
	// 组织节点数

	NodeCount int64 `json:"node_count"`
}

func (o OrgPeer) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "OrgPeer struct{}"
	}

	return strings.Join([]string{"OrgPeer", string(data)}, " ")
}
