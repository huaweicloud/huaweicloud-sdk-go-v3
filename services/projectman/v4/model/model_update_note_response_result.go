package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateNoteResponseResult **参数解释：** 返回结果。 **取值范围：** 不涉及。
type UpdateNoteResponseResult struct {

	// **参数解释：** 返回更新状态。 **取值范围：** success：更新成功。 error：更新失败。
	Status *string `json:"status,omitempty"`
}

func (o UpdateNoteResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateNoteResponseResult struct{}"
	}

	return strings.Join([]string{"UpdateNoteResponseResult", string(data)}, " ")
}
