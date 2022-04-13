package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 节点启停
type NodeAction struct {
	Node *Action `json:"node"`
}

func (o NodeAction) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeAction struct{}"
	}

	return strings.Join([]string{"NodeAction", string(data)}, " ")
}
