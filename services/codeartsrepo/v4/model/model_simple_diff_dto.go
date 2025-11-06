package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleDiffDto struct {

	// **参数解释：** 增加行数。 **取值范围：** 不涉及。
	AddedLine *int32 `json:"added_line,omitempty"`

	// **参数解释：** 删除行数。 **取值范围：** 不涉及。
	DeletedLine *int32 `json:"deleted_line,omitempty"`

	// **参数解释：** 路径。 **取值范围：** 不涉及。
	Path *string `json:"path,omitempty"`
}

func (o SimpleDiffDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleDiffDto struct{}"
	}

	return strings.Join([]string{"SimpleDiffDto", string(data)}, " ")
}
