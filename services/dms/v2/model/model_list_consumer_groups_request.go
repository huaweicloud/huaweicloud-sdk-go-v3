package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListConsumerGroupsRequest struct {
	// 指定的队列ID

	QueueId string `json:"queue_id"`
	// 是否包含死信信息。默认值为：false

	IncludeDeadletter *bool `json:"include_deadletter,omitempty"`
	// 是否查询消费组的详情，默认值为true。  配置为false时，查询的消费组列表不包含消费详情，接口响应速度快。

	IncludeMessagesNum *bool `json:"include_messages_num,omitempty"`
	// 设置每页显示的消费组数量。  page_size和current_page必须同时配置有效值，否则默认查询所有消费组。

	PageSize *int32 `json:"page_size,omitempty"`
	// 设置查询消费组的页码。  page_size和current_page必须同时配置有效值，否则默认查询所有消费组。

	CurrentPage *int32 `json:"current_page,omitempty"`
}

func (o ListConsumerGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumerGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListConsumerGroupsRequest", string(data)}, " ")
}
