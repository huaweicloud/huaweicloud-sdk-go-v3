package model

import (
	"encoding/json"

	"strings"
)

type AuthorizeNa2NodesRequestDto struct {
	// 授权北向NA信息到边缘节点的请求结构体

	NodeIds *interface{} `json:"node_ids"`
}

func (o AuthorizeNa2NodesRequestDto) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AuthorizeNa2NodesRequestDto struct{}"
	}

	return strings.Join([]string{"AuthorizeNa2NodesRequestDto", string(data)}, " ")
}
