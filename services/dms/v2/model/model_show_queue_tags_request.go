package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowQueueTagsRequest struct {
	// 队列ID。

	QueueId string `json:"queue_id"`
}

func (o ShowQueueTagsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowQueueTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowQueueTagsRequest", string(data)}, " ")
}
