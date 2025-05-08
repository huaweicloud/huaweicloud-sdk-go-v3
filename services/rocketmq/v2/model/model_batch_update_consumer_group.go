package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdateConsumerGroup struct {

	// 消费组名称，只能由英文字母、数字、百分号、竖线、中划线、下划线组成，长度3~64个字符。
	Name *string `json:"name,omitempty"`

	// 是否广播。
	Broadcast *bool `json:"broadcast,omitempty"`

	// 最大重试次数，取值范围为1~16。
	RetryMaxTime *int32 `json:"retry_max_time,omitempty"`

	// 是否可以消费。
	Enabled *bool `json:"enabled,omitempty"`

	// 是否按顺序消费（仅RocketMQ实例5.x版本需要填写此参数）。
	ConsumeOrderly *bool `json:"consume_orderly,omitempty"`

	// 消费组描述，长度0~200个字符。
	GroupDesc *string `json:"group_desc,omitempty"`
}

func (o BatchUpdateConsumerGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateConsumerGroup struct{}"
	}

	return strings.Join([]string{"BatchUpdateConsumerGroup", string(data)}, " ")
}
