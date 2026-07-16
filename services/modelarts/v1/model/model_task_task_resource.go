package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TaskTaskResource **参数解释**：训练作业资源规格信息。 **约束限制**：不涉及。
type TaskTaskResource struct {

	// **参数解释**：训练作业选择的资源规格ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	FlavorId *string `json:"flavor_id,omitempty"`

	// **参数解释**：训练作业选择的资源副本数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NodeCount int32 `json:"node_count"`

	// **参数解释**：训练任务选择的资源池ID。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolId *string `json:"pool_id,omitempty"`
}

func (o TaskTaskResource) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TaskTaskResource struct{}"
	}

	return strings.Join([]string{"TaskTaskResource", string(data)}, " ")
}
