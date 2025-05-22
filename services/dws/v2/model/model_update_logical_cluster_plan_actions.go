package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateLogicalClusterPlanActions **参数解释**： 更新逻辑集群增删计划细节信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type UpdateLogicalClusterPlanActions struct {

	// **参数解释**： 逻辑集群增删行动ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 更新逻辑集群增删计划行动类型。 **约束限制**： 不涉及。 **取值范围**： create：创建 delete：删除 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 更新逻辑集群增删计划行为Cron策略表达式。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Strategy *string `json:"strategy,omitempty"`
}

func (o UpdateLogicalClusterPlanActions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLogicalClusterPlanActions struct{}"
	}

	return strings.Join([]string{"UpdateLogicalClusterPlanActions", string(data)}, " ")
}
