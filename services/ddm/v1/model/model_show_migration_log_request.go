package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMigrationLogRequest Request Object
type ShowMigrationLogRequest struct {

	// **参数解释**：  实例ID，此参数是实例的唯一标识。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成，后缀为in09，长度为36个字符。  **默认取值**：  不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**：  分片变更任务 ID。  **约束限制**：  不涉及。  **取值范围**：  只能由英文字母、数字组成。  **默认取值**：  不涉及。
	TaskId string `json:"task_id"`

	// **参数解释**：  分页参数: 起始值。  **约束限制**：  不涉及。  **取值范围**：  大于等于0。  **默认取值**：  默认值是0。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**：  分页参数: 每页记录数。  **约束限制**：  不涉及。  **取值范围**：  大于0且小于等于128。  **默认取值**：  默认值是10。
	Limit *int32 `json:"limit,omitempty"`
}

func (o ShowMigrationLogRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMigrationLogRequest struct{}"
	}

	return strings.Join([]string{"ShowMigrationLogRequest", string(data)}, " ")
}
