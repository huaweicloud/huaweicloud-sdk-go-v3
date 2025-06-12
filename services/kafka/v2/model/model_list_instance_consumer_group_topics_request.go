package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupTopicsRequest Request Object
type ListInstanceConsumerGroupTopicsRequest struct {

	// 引擎。
	Engine string `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组ID。
	Group string `json:"group"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 当次查询返回的最大Topic个数，默认值为10，取值范围为1~50。
	Limit *int32 `json:"limit,omitempty"`

	// 排序规则： - topic：按Topic名称排序。 - partition：按分区数排序。 - messages：按消息数量排序，默认方式。
	SortKey *string `json:"sort_key,omitempty"`

	// 排序方式。 - asc：升序。 - desc：降序，默认方式。
	SortDir *string `json:"sort_dir,omitempty"`

	// Topic名称。
	Topic *string `json:"topic,omitempty"`
}

func (o ListInstanceConsumerGroupTopicsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupTopicsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupTopicsRequest", string(data)}, " ")
}
