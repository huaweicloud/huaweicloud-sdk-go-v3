package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodeAllocatedResourceDto 节点已分配资源详情
type NodeAllocatedResourceDto struct {
	Request *NodeResourceDto `json:"request,omitempty"`

	Limit *NodeResourceDto `json:"limit,omitempty"`
}

func (o NodeAllocatedResourceDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodeAllocatedResourceDto struct{}"
	}

	return strings.Join([]string{"NodeAllocatedResourceDto", string(data)}, " ")
}
