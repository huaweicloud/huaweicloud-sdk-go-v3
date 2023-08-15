package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteQueuePlanRequest Request Object
type DeleteQueuePlanRequest struct {

	// 待删除的队列扩缩容计划的ID
	PlanId int64 `json:"plan_id"`

	// 待删除定时扩缩计划的队列名称
	QueueName string `json:"queue_name"`
}

func (o DeleteQueuePlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteQueuePlanRequest struct{}"
	}

	return strings.Join([]string{"DeleteQueuePlanRequest", string(data)}, " ")
}
