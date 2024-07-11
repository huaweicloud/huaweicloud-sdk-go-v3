package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsumerListOrDetailsRequest Request Object
type ShowConsumerListOrDetailsRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组名称。
	Group string `json:"group"`

	// 待查询的Topic，不指定时查询Topic列表，指定时查询详情。
	Topic *string `json:"topic,omitempty"`

	// 当次查询返回的最大个数，默认值为10，取值范围为1~50。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ShowConsumerListOrDetailsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerListOrDetailsRequest struct{}"
	}

	return strings.Join([]string{"ShowConsumerListOrDetailsRequest", string(data)}, " ")
}
