package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchDeleteQueuePlansRequest Request Object
type BatchDeleteQueuePlansRequest struct {

	// 待删除定时扩缩计划的队列名称
	QueueName string `json:"queue_name"`

	Body *BatchDeleteQueuePlansRequestBody `json:"body,omitempty"`
}

func (o BatchDeleteQueuePlansRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchDeleteQueuePlansRequest struct{}"
	}

	return strings.Join([]string{"BatchDeleteQueuePlansRequest", string(data)}, " ")
}
