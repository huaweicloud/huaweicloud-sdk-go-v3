package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DefectFragmentV2 struct {

	// **参数解释：** 行号。 **取值范围：** 不涉及。
	LineNum *string `json:"line_num,omitempty"`

	// **参数解释：** 行代码内容。 **取值范围：** 不涉及。
	LineContent *string `json:"line_content,omitempty"`

	// **参数解释：** 废弃字段。 **取值范围：** 不涉及。
	StartOffset *int32 `json:"start_offset,omitempty"`

	// **参数解释：** 废弃字段。 **取值范围：** 不涉及。
	EndOffset *int32 `json:"end_offset,omitempty"`
}

func (o DefectFragmentV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefectFragmentV2 struct{}"
	}

	return strings.Join([]string{"DefectFragmentV2", string(data)}, " ")
}
