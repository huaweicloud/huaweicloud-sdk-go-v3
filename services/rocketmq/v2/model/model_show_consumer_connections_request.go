package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsumerConnectionsRequest Request Object
type ShowConsumerConnectionsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组名称。
	Group string `json:"group"`

	// 查询数量，取值范围为1~50。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`

	// 是否查询消费者详细列表，参数为“true”则表示查询详细列表，否则表示查询简易列表。
	IsDetail *bool `json:"is_detail,omitempty"`
}

func (o ShowConsumerConnectionsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerConnectionsRequest struct{}"
	}

	return strings.Join([]string{"ShowConsumerConnectionsRequest", string(data)}, " ")
}
