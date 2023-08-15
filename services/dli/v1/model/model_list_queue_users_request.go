package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueueUsersRequest Request Object
type ListQueueUsersRequest struct {

	// 队列名称。
	QueueName string `json:"queue_name"`
}

func (o ListQueueUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueueUsersRequest struct{}"
	}

	return strings.Join([]string{"ListQueueUsersRequest", string(data)}, " ")
}
