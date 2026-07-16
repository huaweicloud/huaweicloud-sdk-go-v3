package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NodePoolStatusResources **参数解释**：节点池中不同状态的资源量。
type NodePoolStatusResources struct {
	Creating *PoolResourceFlavorCount `json:"creating,omitempty"`

	Available *PoolResourceFlavorCount `json:"available,omitempty"`

	Abnormal *PoolResourceFlavorCount `json:"abnormal,omitempty"`

	Deleting *PoolResourceFlavorCount `json:"deleting,omitempty"`
}

func (o NodePoolStatusResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NodePoolStatusResources struct{}"
	}

	return strings.Join([]string{"NodePoolStatusResources", string(data)}, " ")
}
