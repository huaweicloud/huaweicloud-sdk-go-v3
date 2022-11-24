package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 更新节点配置
type UpdateEdgeNodeBody struct {
	Node *EdgeNodeUpdate `json:"node"`
}

func (o UpdateEdgeNodeBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateEdgeNodeBody struct{}"
	}

	return strings.Join([]string{"UpdateEdgeNodeBody", string(data)}, " ")
}
