package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DefectEvents struct {

	// **参数解释：** 调用链信息。
	Events *[]DefectEvents `json:"events,omitempty"`

	// **参数解释：** 问题详细描述。 **取值范围：** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 修复建议。
	FixSuggestions *[]string `json:"fix_suggestions,omitempty"`

	// **参数解释：** 问题所在文件行号。 **取值范围：** 不涉及。
	Line *int32 `json:"line,omitempty"`

	// **参数解释：** 废弃字段。 **取值范围：** 不涉及。
	EndLine *int32 `json:"end_line,omitempty"`

	// **参数解释：** 废弃字段。 **取值范围：** 不涉及。
	Main *bool `json:"main,omitempty"`

	// **参数解释：** 文件路径。 **取值范围：** 不涉及。
	Path *string `json:"path,omitempty"`

	// **参数解释：** 废弃字段。 **取值范围：** 不涉及。
	Tag *string `json:"tag,omitempty"`

	// **参数解释：** 问题主事件对应的代码片段内容。 **取值范围：** 不涉及。
	MainBuggyCode *string `json:"main_buggy_code,omitempty"`

	// **参数解释：** 代码上下文。 **取值范围：** 不涉及。
	CodeContext *string `json:"code_context,omitempty"`

	// **参数解释：** 代码上下文起始行号。 **取值范围：** 不涉及。
	CodeContextStartLine *int32 `json:"code_context_start_line,omitempty"`
}

func (o DefectEvents) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefectEvents struct{}"
	}

	return strings.Join([]string{"DefectEvents", string(data)}, " ")
}
