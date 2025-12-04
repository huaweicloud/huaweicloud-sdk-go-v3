package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRebuildSlaveStatusResponse Response Object
type ShowRebuildSlaveStatusResponse struct {

	// **参数解释**：  实例ID  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  任务流ID  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	WorkflowId *string `json:"workflow_id,omitempty"`

	// **参数解释**：  上一次重建的时间  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	LastRebuildTime *string `json:"last_rebuild_time,omitempty"`

	// **参数解释**：  下次允许重建的时间  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	NextRebuildTime *string `json:"next_rebuild_time,omitempty"`
	HttpStatusCode  int     `json:"-"`
}

func (o ShowRebuildSlaveStatusResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRebuildSlaveStatusResponse struct{}"
	}

	return strings.Join([]string{"ShowRebuildSlaveStatusResponse", string(data)}, " ")
}
