package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueuePlansRequest Request Object
type ListQueuePlansRequest struct {

	// 待查询定时扩缩计划的队列名称
	QueueName string `json:"queue_name"`
}

func (o ListQueuePlansRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuePlansRequest struct{}"
	}

	return strings.Join([]string{"ListQueuePlansRequest", string(data)}, " ")
}
