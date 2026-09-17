package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateResponseResultProject **参数解释：** 项目。 **取值范围：** 不涉及。
type BatchUpdateResponseResultProject struct {

	// **参数解释：** 项目数字id。 **取值范围：** 不涉及。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 项目uuid。 **取值范围：** 不涉及。
	Identifier *string `json:"identifier,omitempty"`

	// **参数解释：** 批量编辑工作项的总数。 **取值范围：** 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** 项目是否关闭。 **取值范围：** 0（打开） 1（关闭）
	Close *int32 `json:"close,omitempty"`

	// **参数解释：** 批量编辑数量。 **取值范围：** 不涉及。
	Role *int32 `json:"role,omitempty"`

	// **参数解释：** 工作项类型。 **取值范围：** scrum。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 工作项是否归档。 **取值范围：** true(归档) false(未归档)
	Archive *bool `json:"archive,omitempty"`

	// **参数解释：** 项目数量。 **取值范围：** 不涉及。
	MemCount *int32 `json:"mem_count,omitempty"`
}

func (o BatchUpdateResponseResultProject) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateResponseResultProject struct{}"
	}

	return strings.Join([]string{"BatchUpdateResponseResultProject", string(data)}, " ")
}
