package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQueuePropertyRequest Request Object
type CreateQueuePropertyRequest struct {

	// 队列名称
	QueueName string `json:"queue_name"`

	Body *CreateQueuePropertyRequestBody `json:"body,omitempty"`
}

func (o CreateQueuePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueuePropertyRequest struct{}"
	}

	return strings.Join([]string{"CreateQueuePropertyRequest", string(data)}, " ")
}
