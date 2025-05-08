package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupTopicsResponse Response Object
type ListInstanceConsumerGroupTopicsResponse struct {

	// 消费组TOPIC
	Topics *[]GroupTopicEntity `json:"topics,omitempty"`

	// 统计数量
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListInstanceConsumerGroupTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupTopicsResponse", string(data)}, " ")
}
