package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowQueueDetailRequest Request Object
type ShowQueueDetailRequest struct {

	// 指定查询的队列名称。
	QueueName string `json:"queue_name"`
}

func (o ShowQueueDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowQueueDetailRequest", string(data)}, " ")
}
