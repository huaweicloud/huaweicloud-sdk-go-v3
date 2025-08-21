package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowGroupGeneralPolicyResponse Response Object
type ShowGroupGeneralPolicyResponse struct {

	// **参数解释：** 是否禁用fork。
	DisableFork *bool `json:"disable_fork,omitempty"`

	// **参数解释：** 是否禁止开发者创建分支。
	ForbiddenDeveloperCreateBranch *bool `json:"forbidden_developer_create_branch,omitempty"`

	// **参数解释：** 是否禁用开发者创建tag。
	ForbiddenDeveloperCreateTag *bool `json:"forbidden_developer_create_tag,omitempty"`

	// **参数解释：** 禁止开发者创建标签。
	ForbiddenCommitterCreateBranch *bool `json:"forbidden_committer_create_branch,omitempty"`

	// **参数解释：** 分支名称正则表达式。 **取值范围：** 字符串长度不少于1，不超过1000。
	BranchNameRegex *string `json:"branch_name_regex,omitempty"`

	// **参数解释：** 标签名正则表达式。 **取值范围：** 字符串长度不少于1，不超过1000。
	TagNameRegex *string `json:"tag_name_regex,omitempty"`

	// **参数解释：** 生成合并请求预合并。
	GeneratePreMergeRef *bool `json:"generate_pre_merge_ref,omitempty"`

	// **参数解释：** 是否禁止repo访问。
	ForbiddenGitlabAccess *bool `json:"forbidden_gitlab_access,omitempty"`

	// **参数解释：** MR rebase是否禁止触发webhook事件。
	RebaseDisableTriggerWebhook *bool `json:"rebase_disable_trigger_webhook,omitempty"`

	// **参数解释：** 是否开启gpg公钥验证。
	OpenGpgVerified *bool `json:"open_gpg_verified,omitempty"`
	HttpStatusCode  int   `json:"-"`
}

func (o ShowGroupGeneralPolicyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowGroupGeneralPolicyResponse struct{}"
	}

	return strings.Join([]string{"ShowGroupGeneralPolicyResponse", string(data)}, " ")
}
