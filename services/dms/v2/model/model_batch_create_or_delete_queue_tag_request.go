package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateOrDeleteQueueTagRequest struct {

	// 队列ID。
	QueueId string `json:"queue_id" xml:"queue_id"`

	Body *BatchCreateOrDeleteTagReq `json:"body,omitempty" xml:"body"`
}

func (o BatchCreateOrDeleteQueueTagRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateOrDeleteQueueTagRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateOrDeleteQueueTagRequest", string(data)}, " ")
}
