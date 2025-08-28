package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// JobEntities **参数解释**：子任务关联的资源列表
type JobEntities struct {

	// **参数解释**：子任务关联的资源ID。  **取值范围**：不涉及
	ResourceId *string `json:"resource_id,omitempty"`

	// **参数解释**：子任务关联的资源类型。  **取值范围**：不涉及
	ResourceType *string `json:"resource_type,omitempty"`
}

func (o JobEntities) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "JobEntities struct{}"
	}

	return strings.Join([]string{"JobEntities", string(data)}, " ")
}
