package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ItemCommitDto 工作项关联的提交信息
type ItemCommitDto struct {

	// **参数解释：** 关联Id。 **取值范围：** 不涉及
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 合并请求标题。 **取值范围：** 不涉及。
	Title *string `json:"title,omitempty"`

	// **参数解释：** 仓库状态。 **取值范围：** - 1，关联成功。 - 2，关联失败。 - 3，流转成功。
	Result *int32 `json:"result,omitempty"`

	// **参数解释：** 关联提交类型。 **取值范围：** - commit，提交。 - branch，分支。 - mergerequest，合并请求。
	Type *string `json:"type,omitempty"`

	// **参数解释：** 关联失败的失败原因。 **取值范围：** 不涉及。
	Message *string `json:"message,omitempty"`

	// **参数解释：** 工作项Id。 **取值范围：** 不涉及。
	ItemId *string `json:"item_id,omitempty"`

	// **参数解释：** 仓库Id。 **取值范围：** 不涉及。
	RepositoryId *string `json:"repository_id,omitempty"`

	// **参数解释：** 分支名。 **取值范围：** 不涉及。
	BranchName *string `json:"branch_name,omitempty"`

	// **参数解释：** 用户名。 **取值范围：** 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释：** 提交Id。 **取值范围：** 不涉及。
	CommitId *string `json:"commit_id,omitempty"`

	// **参数解释：** 提交短Id。 **取值范围：** 不涉及。
	CommitShortId *string `json:"commit_short_id,omitempty"`

	// **参数解释：** 提交信息。 **取值范围：** 不涉及。
	CommitMsg *string `json:"commit_msg,omitempty"`

	// **参数解释：** 提交访问地址。 **取值范围：** 不涉及。
	CommitUrl *string `json:"commit_url,omitempty"`

	// **参数解释：** iamId。 **取值范围：** 不涉及。
	IamId *string `json:"iam_id,omitempty"`

	// **参数解释：** 创建时间。 **取值范围：** 不涉及。
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释：** 更新时间。 **取值范围：** 不涉及。
	UpdateAt *string `json:"update_at,omitempty"`

	// **参数解释：** 工作项访问地址。 **取值范围：** 不涉及。
	ItemUrl *string `json:"item_url,omitempty"`
}

func (o ItemCommitDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ItemCommitDto struct{}"
	}

	return strings.Join([]string{"ItemCommitDto", string(data)}, " ")
}
