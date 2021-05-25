package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowQueueTagsRequest struct {
	// 队列ID。

	QueueId string `json:"queue_id"`
}

func (o ShowQueueTagsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowQueueTagsRequest struct{}"
	}

	return strings.Join([]string{"ShowQueueTagsRequest", string(data)}, " ")
}
