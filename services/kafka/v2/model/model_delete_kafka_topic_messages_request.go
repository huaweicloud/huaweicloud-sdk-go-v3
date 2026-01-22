package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKafkaTopicMessagesRequest Request Object
type DeleteKafkaTopicMessagesRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// Topic名称。
	Topic string `json:"topic"`

	Body *DeleteKafkaMessageRequestBody `json:"body,omitempty"`
}

func (o DeleteKafkaTopicMessagesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKafkaTopicMessagesRequest struct{}"
	}

	return strings.Join([]string{"DeleteKafkaTopicMessagesRequest", string(data)}, " ")
}
