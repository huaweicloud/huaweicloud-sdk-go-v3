package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

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

	// 消费组描述，长度0~200个字符。
	GroupDesc *string `json:"group_desc,omitempty"`

	// 最大重试次数，取值范围为1~16。
	RetryMaxTime *int32 `json:"retry_max_time,omitempty"`

	// 创建时间戳。
	CreatedAt *int64 `json:"createdAt,omitempty"`

	// 权限集。
	Permissions *[]string `json:"permissions,omitempty"`

	// 是否按顺序消费。
	ConsumeOrderly *bool `json:"consume_orderly,omitempty"`
}

func (o ConsumerGroup) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerGroup struct{}"
	}

	return strings.Join([]string{"ConsumerGroup", string(data)}, " ")
}
