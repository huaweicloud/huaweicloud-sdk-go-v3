package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupMessageOffsetResponse Response Object
type ListInstanceConsumerGroupMessageOffsetResponse struct {

	// 消费组消息位点详情
	GroupMessageOffsets *[]GroupMessageOffsetsDetailEntity `json:"group_message_offsets,omitempty"`

	// 总数
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceConsumerGroupMessageOffsetResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupMessageOffsetResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupMessageOffsetResponse", string(data)}, " ")
}
