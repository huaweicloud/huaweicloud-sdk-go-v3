package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeQueuePlanRequest Request Object
type ChangeQueuePlanRequest struct {

	// 待修改的队列扩缩容计划的ID
	PlanId string `json:"plan_id"`

	// 待删除定时扩缩计划的队列名称
	QueueName string `json:"queue_name"`

	Body *SetQueuePlanReq `json:"body,omitempty"`
}

func (o ChangeQueuePlanRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeQueuePlanRequest struct{}"
	}

	return strings.Join([]string{"ChangeQueuePlanRequest", string(data)}, " ")
}
