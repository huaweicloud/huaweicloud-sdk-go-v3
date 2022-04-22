package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type DeleteQueueRequest struct {

	// 待删除的队列ID
	QueueId string `json:"queue_id"`
}

func (o DeleteQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueueRequest struct{}"
	}

	return strings.Join([]string{"DeleteQueueRequest", string(data)}, " ")
}
