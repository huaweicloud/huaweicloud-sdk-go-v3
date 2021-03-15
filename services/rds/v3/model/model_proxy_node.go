package model

import (
	"encoding/json"

	"strings"
)

type ProxyNode struct {
	// Proxy节点ID。
	Id string `json:"id"`
	// Proxy节点名称。
	Name string `json:"name"`
	// Proxy节点角色。 包括master和slave。
	Role string `json:"role"`
	// 可用区。
	AzCode string `json:"az_code"`
	// Proxy节点状态。
	Status string `json:"status"`
	// Proxy节点是否被冻结。 0表示未冻结，1表示冻结，2表示冻结删除。
	FrozenFlag int32 `json:"frozen_flag"`
}

func (o ProxyNode) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ProxyNode struct{}"
	}

	return strings.Join([]string{"ProxyNode", string(data)}, " ")
}
