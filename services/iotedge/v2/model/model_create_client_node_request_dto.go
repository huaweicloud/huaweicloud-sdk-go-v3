package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateClientNodeRequestDto 分配推送路由到边缘节点请求结构体
type CreateClientNodeRequestDto struct {

	// 客户端节点ID，即边缘节点ID
	ClientNodeId string `json:"client_node_id"`
}

func (o CreateClientNodeRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateClientNodeRequestDto struct{}"
	}

	return strings.Join([]string{"CreateClientNodeRequestDto", string(data)}, " ")
}
