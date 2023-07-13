package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupsRequest Request Object
type ListInstanceConsumerGroupsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组名称。
	Group *string `json:"group,omitempty"`

	// 查询数量。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListInstanceConsumerGroupsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupsRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupsRequest", string(data)}, " ")
}
