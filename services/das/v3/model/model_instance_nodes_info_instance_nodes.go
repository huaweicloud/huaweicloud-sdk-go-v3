package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstanceNodesInfoInstanceNodes struct {

	// 节点ID
	Id *string `json:"id,omitempty"`

	// 节点名
	Name *string `json:"name,omitempty"`

	// 节点角色
	Role *string `json:"role,omitempty"`

	// 节点状态
	Status *string `json:"status,omitempty"`

	// 节点类型
	Type *string `json:"type,omitempty"`
}

func (o InstanceNodesInfoInstanceNodes) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceNodesInfoInstanceNodes struct{}"
	}

	return strings.Join([]string{"InstanceNodesInfoInstanceNodes", string(data)}, " ")
}
