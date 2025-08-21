package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowDiffCommitResponse Response Object
type ShowDiffCommitResponse struct {

	// **参数解释：** 差异信息。 **取值范围：** 不涉及。
	Diffs *[]DiffNoLineDto `json:"diffs,omitempty"`

	DiffRefs *DiffRefsDto `json:"diff_refs,omitempty"`

	// **参数解释：** 增加行数。 **取值范围：** 不涉及。
	AddedLines *int32 `json:"added_lines,omitempty"`

	// **参数解释：** 删除行数。 **取值范围：** 不涉及。
	RemovedLines *int32 `json:"removed_lines,omitempty"`

	// **参数解释：** 更改文件数目。 **取值范围：** 不涉及。
	ChangeFileCount *int32 `json:"change_file_count,omitempty"`

	// **参数解释：** 更改行数。 **取值范围：** 不涉及。
	ChangeLineCount *int32 `json:"change_line_count,omitempty"`

	// **参数解释：** 是否过大。 **取值范围：** - true，大文件。 - false，非大文件
	TooLarge       *bool `json:"too_large,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowDiffCommitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowDiffCommitResponse struct{}"
	}

	return strings.Join([]string{"ShowDiffCommitResponse", string(data)}, " ")
}
