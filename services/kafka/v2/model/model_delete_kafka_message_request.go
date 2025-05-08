package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKafkaMessageRequest Request Object
type DeleteKafkaMessageRequest struct {

	// 实例ID。
	InstanceId string `json:"instance_id"`

	// Topic名称。
	Topic string `json:"topic"`

	Body *DeleteKafkaMessageRequestBody `json:"body,omitempty"`
}

func (o DeleteKafkaMessageRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKafkaMessageRequest struct{}"
	}

	return strings.Join([]string{"DeleteKafkaMessageRequest", string(data)}, " ")
}
