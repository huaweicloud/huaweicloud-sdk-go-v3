package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommitDiffMetadataResponse Response Object
type ShowCommitDiffMetadataResponse struct {

	// **参数解释：** 差异信息。 **取值范围：** 不涉及。
	Diffs *[]DiffDto `json:"diffs,omitempty"`

	DiffRefs *DiffRefsDto `json:"diff_refs,omitempty"`

	// **参数解释：** 增加行数。
	AddedLines *int32 `json:"added_lines,omitempty"`

	// **参数解释：** 删除行数。
	RemovedLines *int32 `json:"removed_lines,omitempty"`

	// **参数解释：** 修改文件数量。
	ChangeFileCount *int32 `json:"change_file_count,omitempty"`

	// **参数解释：** 改变行数数量。
	ChangeLineCount *int32 `json:"change_line_count,omitempty"`

	// **参数解释：** 是否为大文件。 **取值范围：** - true，大文件。 - false，非大文件
	TooLarge       *bool `json:"too_large,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowCommitDiffMetadataResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommitDiffMetadataResponse struct{}"
	}

	return strings.Join([]string{"ShowCommitDiffMetadataResponse", string(data)}, " ")
}
