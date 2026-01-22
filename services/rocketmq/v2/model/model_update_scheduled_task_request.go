package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateScheduledTaskRequest Request Object
type UpdateScheduledTaskRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 定时任务ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TaskId string `json:"task_id"`

	// **参数解释**： 修改定时任务的执行时间。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ExecuteAt *string `json:"execute_at,omitempty"`

	// **参数解释**： 修改定时任务状态。 **约束限制**： 不涉及。 **取值范围**： - CANCELLED。 **默认取值**： 不涉及。
	Status *string `json:"status,omitempty"`
}

func (o UpdateScheduledTaskRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateScheduledTaskRequest struct{}"
	}

	return strings.Join([]string{"UpdateScheduledTaskRequest", string(data)}, " ")
}
