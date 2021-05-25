package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type BatchCreateOrDeleteQueueTagRequest struct {
	// 队列ID。

	QueueId string `json:"queue_id"`

	Body *BatchCreateOrDeleteTagReq `json:"body,omitempty"`
}

func (o BatchCreateOrDeleteQueueTagRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteQueueTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteQueueTagRequest", string(data)}, " ")
}
