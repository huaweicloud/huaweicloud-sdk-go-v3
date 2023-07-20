package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueueRequest Request Object
type DeleteQueueRequest struct {

	// 指定删除的队列名称。
	QueueName string `json:"queue_name"`
}

func (o DeleteQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueueRequest struct{}"
	}

	return strings.Join([]string{"DeleteQueueRequest", string(data)}, " ")
}
