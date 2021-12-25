package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘节点参数
type Node struct {
	Node *EdgeNode `json:"node"`
}

func (o Node) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Node struct{}"
	}

	return strings.Join([]string{"Node", string(data)}, " ")
}
