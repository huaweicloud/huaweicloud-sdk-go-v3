package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateConsumerGroupRequest Request Object
type UpdateConsumerGroupRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// 消费组名称。
	Group string `json:"group"`

	Body *ConsumerGroup `json:"body,omitempty"`
}

func (o UpdateConsumerGroupRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateConsumerGroupRequest struct{}"
	}

	return strings.Join([]string{"UpdateConsumerGroupRequest", string(data)}, " ")
}
