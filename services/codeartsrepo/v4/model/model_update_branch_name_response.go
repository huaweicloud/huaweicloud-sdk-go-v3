package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBranchNameResponse Response Object
type UpdateBranchNameResponse struct {

	// **参数解释：** 旧分支名称。  **约束限制：** 不支持以 - . refs/heads/ refs/remotes/ 开头，不支持空格 [ \\ < ~ ^ : ? * ! ( ) ' \" | 等特殊字符，不支持以. / .lock结尾。  **取值范围：** 字符串长度不少于1，不超过200。 **默认取值：** 不涉及。
	OldBranchName *string `json:"old_branch_name,omitempty"`

	// **参数解释：** 旧分支最新提交ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	OldBranchCommitId *string `json:"old_branch_commit_id,omitempty"`

	// **参数解释：** 新分支名称。  **约束限制：** 不支持以 - . refs/heads/ refs/remotes/ 开头，不支持空格 [ \\ < ~ ^ : ? * ! ( ) ' \" | 等特殊字符，不支持以. / .lock结尾。  **取值范围：** 字符串长度不少于1，不超过200。 **默认取值：** 不涉及。
	NewBranchName *string `json:"new_branch_name,omitempty"`

	// **参数解释：** 新分支最新提交ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	NewBranchCommitId *string `json:"new_branch_commit_id,omitempty"`
	HttpStatusCode    int     `json:"-"`
}

func (o UpdateBranchNameResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBranchNameResponse struct{}"
	}

	return strings.Join([]string{"UpdateBranchNameResponse", string(data)}, " ")
}
