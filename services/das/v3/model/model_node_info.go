package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeInfo 实例节点信息
type NodeInfo struct {

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

func (o NodeInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeInfo struct{}"
	}

	return strings.Join([]string{"NodeInfo", string(data)}, " ")
}
