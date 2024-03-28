package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateQueuePlanRequest Request Object
type CreateQueuePlanRequest struct {

	// 需要设置定时扩缩计划的队列名称，名称只能包含数字、英文字母和下划线，但不能是纯数字，且不能以下划线开头。
	QueueName string `json:"queue_name"`

	Body *QueuePlanRequestBody `json:"body,omitempty"`
}

func (o CreateQueuePlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueuePlanRequest struct{}"
	}

	return strings.Join([]string{"CreateQueuePlanRequest", string(data)}, " ")
}
