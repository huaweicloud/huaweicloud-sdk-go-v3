package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInstanceConsumerGroupMembersRequest Request Object
type ListInstanceConsumerGroupMembersRequest struct {

	// 引擎。
	Engine string `json:"engine"`

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组ID。
	Group string `json:"group"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 当次查询返回的最大消费组成员个数，默认值为10，取值范围为1~50。
	Limit *int32 `json:"limit,omitempty"`

	// 消费者地址。
	Host *string `json:"host,omitempty"`

	// 消费者ID。
	MemberId *string `json:"member_id,omitempty"`
}

func (o ListInstanceConsumerGroupMembersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceConsumerGroupMembersRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceConsumerGroupMembersRequest", string(data)}, " ")
}
