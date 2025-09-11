package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RecordInfoResponseResult struct {

	// **参数解释**: 主键id。 **取值范围**: 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**: 动作。 **取值范围**: 不涉及。
	Action *string `json:"action,omitempty"`

	// **参数解释**: 操作状态。 **取值范围**: 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**: 信息。 **取值范围**: 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释**: 实体id。 **取值范围**: 不涉及。
	EntityId *string `json:"entity_id,omitempty"`

	// **参数解释**: 实体类型。 **取值范围**: 不涉及。
	EntityType *string `json:"entity_type,omitempty"`

	// **参数解释**: 工作流id。 **取值范围**: 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// **参数解释**: 实例id。 **取值范围**: 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**: 创建时间。 **取值范围**: 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**: 更新时间。 **取值范围**: 不涉及。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释**: 扩展信息。 **取值范围**: 不涉及。
	ExtendedInfo *interface{} `json:"extended_info,omitempty"`
}

func (o RecordInfoResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecordInfoResponseResult struct{}"
	}

	return strings.Join([]string{"RecordInfoResponseResult", string(data)}, " ")
}
