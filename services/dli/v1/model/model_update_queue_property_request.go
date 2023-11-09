package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQueuePropertyRequest Request Object
type UpdateQueuePropertyRequest struct {

	// 队列名称
	QueueName string `json:"queue_name"`

	Body *UpdateQueuePropertyRequestBody `json:"body,omitempty"`
}

func (o UpdateQueuePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQueuePropertyRequest struct{}"
	}

	return strings.Join([]string{"UpdateQueuePropertyRequest", string(data)}, " ")
}
