package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowTopicStatusRespQueues struct {

	// 队列ID。
	Id *int32 `json:"id,omitempty"`

	// 最小偏移量。
	MinOffset *int32 `json:"min_offset,omitempty"`

	// 最大偏移量。
	MaxOffset *int32 `json:"max_offset,omitempty"`

	// 最后一条消息的时间。
	LastMessageTime *int64 `json:"last_message_time,omitempty"`
}

func (o ShowTopicStatusRespQueues) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicStatusRespQueues struct{}"
	}

	return strings.Join([]string{"ShowTopicStatusRespQueues", string(data)}, " ")
}
