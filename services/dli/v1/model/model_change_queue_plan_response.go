package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ChangeQueuePlanResponse Response Object
type ChangeQueuePlanResponse struct {

	// 定时扩缩容计划对应的队列名称
	QueueName *string `json:"queue_name,omitempty"`

	// 扩缩容计划的ID编号
	PlanId *string `json:"plan_id,omitempty"`

	// 请求执行是否成功。“true”表示请求执行成功。
	IsSuccess *bool `json:"is_success,omitempty"`

	// 系统提示信息，执行成功时，信息可能为空。
	Message        *string `json:"message,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ChangeQueuePlanResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChangeQueuePlanResponse struct{}"
	}

	return strings.Join([]string{"ChangeQueuePlanResponse", string(data)}, " ")
}
