package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WarmPoolInfo struct {

	// 暖池最小容量
	MinCapacity *int32 `json:"min_capacity,omitempty"`

	// 暖池最大容量
	MaxCapacity *int32 `json:"max_capacity,omitempty"`

	// 实例初始化等待时间
	InstanceInitWaitTime *int32 `json:"instance_init_wait_time,omitempty"`

	// 暖池状态
	Status *string `json:"status,omitempty"`
}

func (o WarmPoolInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WarmPoolInfo struct{}"
	}

	return strings.Join([]string{"WarmPoolInfo", string(data)}, " ")
}
