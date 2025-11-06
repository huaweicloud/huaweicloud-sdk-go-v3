package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateRepositoryGeneralPolicyResponse Response Object
type UpdateRepositoryGeneralPolicyResponse struct {

	// **参数解释：** 是否禁止fork该仓库。 **约束限制：** 不涉及。 **取值范围：** - true，禁止fork。 - false，允许fork。 **默认取值：** 不涉及。
	DisableFork *bool `json:"disable_fork,omitempty"`

	// **参数解释：** 是否预合并MR。 **约束限制：** 不涉及。 **取值范围：** - true，禁止预合并MR。 - false，允许预合并MR。 **默认取值：** 不涉及。
	GeneratePreMergeRef *bool `json:"generate_pre_merge_ref,omitempty"`

	// **参数解释：** 分支名规则。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	BranchNameRegex *string `json:"branch_name_regex,omitempty"`

	// **参数解释：** Tag名规则。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	TagNameRegex *string `json:"tag_name_regex,omitempty"`

	// **参数解释：** 是否禁止开发者创建分支。 **约束限制：** 不涉及。 **取值范围：** - true，禁止开发者创建分支。 - false，允许开发者创建分支。
	ForbiddenDeveloperCreateBranch *bool `json:"forbidden_developer_create_branch,omitempty"`

	// **参数解释：** 开发人员创建分支权限白名单。
	CreateBranchWhitelistUsers *[]PushRuleDevelopersDto `json:"create_branch_whitelist_users,omitempty"`
	HttpStatusCode             int                      `json:"-"`
}

func (o UpdateRepositoryGeneralPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateRepositoryGeneralPolicyResponse struct{}"
	}

	return strings.Join([]string{"UpdateRepositoryGeneralPolicyResponse", string(data)}, " ")
}
