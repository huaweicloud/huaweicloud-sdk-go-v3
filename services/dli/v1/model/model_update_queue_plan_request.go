package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateQueuePlanRequest Request Object
type UpdateQueuePlanRequest struct {

	// 待修改的队列扩缩容计划的ID
	PlanId string `json:"plan_id"`

	// 待删除定时扩缩计划的队列名称
	QueueName string `json:"queue_name"`

	Body *QueuePlanRequestBody `json:"body,omitempty"`
}

func (o UpdateQueuePlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateQueuePlanRequest struct{}"
	}

	return strings.Join([]string{"UpdateQueuePlanRequest", string(data)}, " ")
}
