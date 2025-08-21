package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitParams **参数解释：** 提交信息参数。
type CommitParams struct {

	// **参数解释：** 分支名称。 **取值范围：** 不涉及。
	Branch string `json:"branch"`

	// **参数解释：** 提交信息。 **取值范围：** 不涉及。
	CommitMessage string `json:"commit_message"`

	// **参数解释：** 在提交时执行的操作。 **取值范围：** 不涉及。
	Actions []ActionDto `json:"actions"`

	// **参数解释：** 从该分支开始新的提交。  **取值范围：** 不涉及。
	StartBranch *string `json:"start_branch,omitempty"`

	// **参数解释：** 作者邮箱。  **取值范围：** 不涉及。
	AuthorEmail *string `json:"author_email,omitempty"`

	// **参数解释：** 作者名称。  **取值范围：** 不涉及。
	AuthorName *string `json:"author_name,omitempty"`

	// **参数解释：** 是否包括提交统计信息。 **取值范围：** - true，包括提交统计信息。 - false，不包括提交统计信息
	Stats *bool `json:"stats,omitempty"`

	// **参数解释：** 是否强制提交。 **取值范围：** - true，强制提交。 - false，非强制提交
	Force *bool `json:"force,omitempty"`
}

func (o CommitParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitParams struct{}"
	}

	return strings.Join([]string{"CommitParams", string(data)}, " ")
}
