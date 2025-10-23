package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RenameFileBodyDto struct {

	// **参数解释：** 文件路径 **取值范围：** 不涉及。
	FilePath string `json:"file_path"`

	// **参数解释：** 分支名 **取值范围：** 不涉及。
	BranchName string `json:"branch_name"`

	// **参数解释：** 提交信息 **取值范围：** 不涉及。
	CommitMessage string `json:"commit_message"`

	// **参数解释：** 基于分支名（即基于某分支，向其他分支进行推送） **取值范围：** 不涉及。
	StartBranch *string `json:"start_branch,omitempty"`

	// **参数解释：** 作者邮箱 **取值范围：** 不涉及。
	AuthorEmail *string `json:"author_email,omitempty"`

	// **参数解释：** 作者名称 **取值范围：** 不涉及。
	AuthorName *string `json:"author_name,omitempty"`

	// **参数解释：** 改名前地址 **取值范围：** 不涉及。
	PreviousPath string `json:"previous_path"`

	// **参数解释：** 是否迁移内容（与content字段不能同时为空） **取值范围：** 不涉及。
	InferContent *bool `json:"infer_content,omitempty"`

	// **参数解释：** 文件内容（与infer_content字段不能同时为空） **取值范围：** 不涉及。
	Content *string `json:"content,omitempty"`

	// **参数解释：** 编码方式。 **取值范围：** - text。 - base64.
	Encoding *string `json:"encoding,omitempty"`

	// **参数解释：** 上次已知的文件提交ID。 **取值范围：** 不涉及。
	LastCommitId *string `json:"last_commit_id,omitempty"`
}

func (o RenameFileBodyDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RenameFileBodyDto struct{}"
	}

	return strings.Join([]string{"RenameFileBodyDto", string(data)}, " ")
}
