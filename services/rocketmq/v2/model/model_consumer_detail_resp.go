package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerDetailResp struct {

	// 消费堆积总数。
	Lag *int64 `json:"lag,omitempty"`

	// 消息总数。
	MaxOffset *int64 `json:"max_offset,omitempty"`

	// 已消费消息数。
	ConsumerOffset *int64 `json:"consumer_offset,omitempty"`

	// Topic关联代理（当查询Topic消费“详情”才显示此参数）。
	Brokers *[]Brokers `json:"brokers,omitempty"`
}

func (o ConsumerDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerDetailResp struct{}"
	}

	return strings.Join([]string{"ConsumerDetailResp", string(data)}, " ")
}
