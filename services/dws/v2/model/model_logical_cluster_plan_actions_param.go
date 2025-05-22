package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LogicalClusterPlanActionsParam **参数解释**： 用于提交逻辑集群增删计划的行为信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type LogicalClusterPlanActionsParam struct {

	// **参数解释**： 定时增删计划行为类型，取值范围为（create|delete）。 **约束限制**： 不涉及。 **取值范围**： create：创建 delete：删除 **默认取值**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 周期性定时增删计划，Cron策略表达式：如\"0 0 0 ? * 3\"。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Strategy *string `json:"strategy,omitempty"`
}

func (o LogicalClusterPlanActionsParam) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LogicalClusterPlanActionsParam struct{}"
	}

	return strings.Join([]string{"LogicalClusterPlanActionsParam", string(data)}, " ")
}
