package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerGroup struct {

	// 是否可以消费。
	Enabled *bool `json:"enabled,omitempty"`

	// 是否广播。
	Broadcast *bool `json:"broadcast,omitempty"`

	// 关联的代理列表。
	Brokers *[]string `json:"brokers,omitempty"`

	// 消费组名称，只能由英文字母、数字、百分号、竖线、中划线、下划线组成，长度3~64个字符。
	Name *string `json:"name,omitempty"`

	// 最大重试次数。
	RetryMaxTime float32 `json:"retry_max_time,omitempty"`

	// 是否重头消费。
	FromBeginning *bool `json:"from_beginning,omitempty"`
}

func (o ConsumerGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerGroup struct{}"
	}

	return strings.Join([]string{"ConsumerGroup", string(data)}, " ")
}
