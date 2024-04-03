package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueuePropertyRequest Request Object
type DeleteQueuePropertyRequest struct {

	// 队列名称
	QueueName string `json:"queue_name"`

	Body *DeleteQueuePropertyRequestBody `json:"body,omitempty"`
}

func (o DeleteQueuePropertyRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueuePropertyRequest struct{}"
	}

	return strings.Join([]string{"DeleteQueuePropertyRequest", string(data)}, " ")
}
