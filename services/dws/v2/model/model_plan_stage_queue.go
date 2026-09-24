package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlanStageQueue **参数解释**： 资源池信息。 **取值范围**： 不涉及。
type PlanStageQueue struct {

	// **参数解释**： 资源池名称。 **取值范围**： 不涉及。
	QueueName string `json:"queue_name"`

	// **参数解释**： 资源配置队列。 **取值范围**： 不涉及。
	QueueResources []WorkloadResourceItem `json:"queue_resources"`
}

func (o PlanStageQueue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlanStageQueue struct{}"
	}

	return strings.Join([]string{"PlanStageQueue", string(data)}, " ")
}
