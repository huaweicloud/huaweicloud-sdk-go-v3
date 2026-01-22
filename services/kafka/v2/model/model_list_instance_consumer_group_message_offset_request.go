package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupMessageOffsetRequest Request Object
type ListInstanceConsumerGroupMessageOffsetRequest struct {

	// 消息引擎。
	Engine string `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组名称。
	Group string `json:"group"`

	// topic名称。查询消费组消息位点时必填项。
	Topic string `json:"topic"`

	// 分区名称。
	Partition *string `json:"partition,omitempty"`

	// 偏移值。
	Offset *string `json:"offset,omitempty"`

	// 最大值。
	Limit *string `json:"limit,omitempty"`
}

func (o ListInstanceConsumerGroupMessageOffsetRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupMessageOffsetRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupMessageOffsetRequest", string(data)}, " ")
}
