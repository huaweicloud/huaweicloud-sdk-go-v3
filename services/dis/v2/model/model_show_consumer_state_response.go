package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsumerStateResponse Response Object
type ShowConsumerStateResponse struct {

	// 是否还有更多满足条件的App。  - true：是。 - false：否。
	HasMore *bool `json:"has_more,omitempty"`

	// 要查询的通道名称。
	StreamName *string `json:"stream_name,omitempty"`

	// 要查询的APP的名称，用户数据消费程序的唯一标识符。
	AppName *string `json:"app_name,omitempty"`

	// 当前分区消费状态.
	PartitionConsumingStates *[]PartitionConsumingStates `json:"partition_consuming_states,omitempty"`
	HttpStatusCode           int                         `json:"-"`
}

func (o ShowConsumerStateResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerStateResponse struct{}"
	}

	return strings.Join([]string{"ShowConsumerStateResponse", string(data)}, " ")
}
