package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQueueRequest Request Object
type ShowQueueRequest struct {

	// 指定查询的队列名称。
	QueueName string `json:"queue_name"`
}

func (o ShowQueueRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueRequest struct{}"
	}

	return strings.Join([]string{"ShowQueueRequest", string(data)}, " ")
}
