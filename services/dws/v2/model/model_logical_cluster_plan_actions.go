package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterPlanActions **参数解释**： 逻辑集群增删计划任务信息。 **取值范围**： 不涉及。
type LogicalClusterPlanActions struct {

	// **参数解释**： 下一次触发时间。 **取值范围**： 不涉及。
	NextFireTime *string `json:"next_fire_time,omitempty"`

	// **参数解释**： 失败原因。 **取值范围**： 不涉及。
	FailedReason *string `json:"failed_reason,omitempty"`

	// **参数解释**： 任务ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 类型。 **取值范围**： create：创建 delete：删除
	Type *string `json:"type,omitempty"`

	// **参数解释**： 周期信息。Cron表达式：如\"0 0 0 ? * 3\"。 **取值范围**： 不涉及。
	Strategy *string `json:"strategy,omitempty"`

	// 行动状态，取值范围为(running|waiting|deleted|finished|disabled|failed)
	Status *string `json:"status,omitempty"`

	// **参数解释**： 上一次触发时间。 **取值范围**： 不涉及。
	PreFireTime *string `json:"pre_fire_time,omitempty"`
}

func (o LogicalClusterPlanActions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterPlanActions struct{}"
	}

	return strings.Join([]string{"LogicalClusterPlanActions", string(data)}, " ")
}
