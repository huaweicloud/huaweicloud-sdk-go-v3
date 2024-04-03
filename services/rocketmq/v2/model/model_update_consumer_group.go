package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateConsumerGroup struct {

	// 是否可以消费。
	Enabled bool `json:"enabled"`

	// 是否广播。
	Broadcast bool `json:"broadcast"`

	// 关联的代理列表。
	Brokers *[]string `json:"brokers,omitempty"`

	// 待修改参数的消费组（消费组名称不支持修改）。
	Name *string `json:"name,omitempty"`

	// 最大重试次数，取值范围为1~16。
	RetryMaxTime int32 `json:"retry_max_time"`
}

func (o UpdateConsumerGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConsumerGroup struct{}"
	}

	return strings.Join([]string{"UpdateConsumerGroup", string(data)}, " ")
}
