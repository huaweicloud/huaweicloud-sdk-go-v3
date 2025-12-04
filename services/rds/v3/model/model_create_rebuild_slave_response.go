package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateRebuildSlaveResponse Response Object
type CreateRebuildSlaveResponse struct {

	// **参数解释**：  实例ID  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**：  任务流ID  **约束限制**：  不涉及。  **取值范围**：  不涉及。  **默认取值**：  不涉及。
	WorkflowId     *string `json:"workflow_id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateRebuildSlaveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRebuildSlaveResponse struct{}"
	}

	return strings.Join([]string{"CreateRebuildSlaveResponse", string(data)}, " ")
}
