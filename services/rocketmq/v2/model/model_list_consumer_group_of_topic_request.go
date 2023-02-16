package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListConsumerGroupOfTopicRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 主题名称。
	Topic string `json:"topic"`

	// 当次查询返回的最大个数,默认值为10,取值范围为1~50。
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量，表示从此偏移量开始查询， offset大于等于0。
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListConsumerGroupOfTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumerGroupOfTopicRequest struct{}"
	}

	return strings.Join([]string{"ListConsumerGroupOfTopicRequest", string(data)}, " ")
}
