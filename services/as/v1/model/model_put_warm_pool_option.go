package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PutWarmPoolOption struct {

	// 暖池最小容量，暖池最小与最大容量必须同时填写或不填写。暖池最小容量和最大容量填写时必须相等
	MinCapacity *int32 `json:"min_capacity,omitempty"`

	// 暖池最大容量，暖池最小与最大容量必须同时填写或不填写。暖池最小容量和最大容量填写时必须相等
	MaxCapacity *int32 `json:"max_capacity,omitempty"`

	// 实例初始化等待时间，单位：秒
	InstanceInitWaitTime *int32 `json:"instance_init_wait_time,omitempty"`
}

func (o PutWarmPoolOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutWarmPoolOption struct{}"
	}

	return strings.Join([]string{"PutWarmPoolOption", string(data)}, " ")
}
