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
}

func (o ListConsumerGroupOfTopicRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListConsumerGroupOfTopicRequest struct{}"
	}

	return strings.Join([]string{"ListConsumerGroupOfTopicRequest", string(data)}, " ")
}
