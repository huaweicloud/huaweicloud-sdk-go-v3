package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DefEntryDto **参数解释：** 代码导航def相关信息
type DefEntryDto struct {

	// **参数解释：** 标记名称。 **约束限制：** 不涉及。
	TagName *string `json:"tag_name,omitempty"`

	// **参数解释：** 文件路径。 **约束限制：** 不涉及。
	FilePath *string `json:"file_path,omitempty"`

	// **参数解释：** blob文件ID。 **约束限制：** 不涉及。
	Blob *string `json:"blob,omitempty"`

	// **参数解释：** 索引行简要内容。 **约束限制：** 不涉及。
	LineImage *string `json:"line_image,omitempty"`

	// **参数解释：** 行号。 **约束限制：** 不涉及。
	LineNumber *int32 `json:"line_number,omitempty"`

	// **参数解释：** 范围信息。 **约束限制：** 不涉及。
	Range *string `json:"range,omitempty"`

	// **参数解释：** 语法类型。 **约束限制：** 不涉及。
	SyntaxType *string `json:"syntax_type,omitempty"`

	// **参数解释：** 所在版本号（commit id）。 **约束限制：** 不涉及。
	Revision *string `json:"revision,omitempty"`

	// **参数解释：** 其他信息。 **约束限制：** 不涉及。
	Extend *string `json:"extend,omitempty"`
}

func (o DefEntryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DefEntryDto struct{}"
	}

	return strings.Join([]string{"DefEntryDto", string(data)}, " ")
}
