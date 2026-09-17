package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchUpdateResponseResult **参数解释：** 返回结果。 **取值范围：** 不涉及。
type BatchUpdateResponseResult struct {
	Project *BatchUpdateResponseResultProject `json:"project,omitempty"`

	// **参数解释：** 历史记录id。 **取值范围：** 不涉及。
	JournalIds *[]string `json:"journal_ids,omitempty"`

	// **参数解释：** 编辑失败的工作项。 **取值范围：** 不涉及。
	ErrorIssues *[]int32 `json:"error_issues,omitempty"`

	// **参数解释：** 工作项的迭代版本。 **取值范围：** 不涉及。
	VersionsIssues *[]string `json:"versions_issues,omitempty"`

	// **参数解释：** 编辑成功的工作项。 **取值范围：** 不涉及。
	SuccessIssues *[]string `json:"success_issues,omitempty"`
}

func (o BatchUpdateResponseResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdateResponseResult struct{}"
	}

	return strings.Join([]string{"BatchUpdateResponseResult", string(data)}, " ")
}
