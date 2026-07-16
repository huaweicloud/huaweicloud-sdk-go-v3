package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PoolStatusResources **参数解释**：资源池中不同状态的资源信息。
type PoolStatusResources struct {
	Creating *PoolResourceFlavorCount `json:"creating,omitempty"`

	Available *PoolResourceFlavorCount `json:"available,omitempty"`

	Abnormal *PoolResourceFlavorCount `json:"abnormal,omitempty"`

	Deleting *PoolResourceFlavorCount `json:"deleting,omitempty"`
}

func (o PoolStatusResources) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolStatusResources struct{}"
	}

	return strings.Join([]string{"PoolStatusResources", string(data)}, " ")
}
