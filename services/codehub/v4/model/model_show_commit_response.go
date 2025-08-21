package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowCommitResponse Response Object
type ShowCommitResponse struct {

	// id
	Id *string `json:"id,omitempty"`

	// 消息
	Message *string `json:"message,omitempty"`

	// 父节点提交ID
	ParentIds *[]string `json:"parent_ids,omitempty"`

	// 创建该分支的时间
	AuthoredDate *string `json:"authored_date,omitempty"`

	// 创建者名称
	AuthorName *string `json:"author_name,omitempty"`

	// 创建者邮箱
	AuthorEmail *string `json:"author_email,omitempty"`

	// 代码提交的日期和时间
	CommittedDate *string `json:"committed_date,omitempty"`

	// 提交者名称
	CommitterName *string `json:"committer_name,omitempty"`

	// 提交者邮箱
	CommitterEmail *string `json:"committer_email,omitempty"`

	// 是否打开gpg校验
	OpenGpgVerified *bool `json:"open_gpg_verified,omitempty"`

	// 验证状态
	VerificationStatus *int32 `json:"verification_status,omitempty"`

	// GPG公钥的标识符
	GpgPrimaryKeyId *string `json:"gpg_primary_key_id,omitempty"`

	// 用户名
	Name *string `json:"name,omitempty"`

	// 昵称
	GpgNickName *string `json:"gpg_nick_name,omitempty"`

	// 租户名
	GpgTenantName *string `json:"gpg_tenant_name,omitempty"`

	// 特定GPG用户相关的信息
	GpgUserName *string `json:"gpg_user_name,omitempty"`

	// 短id
	ShortId *string `json:"short_id,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// author_avatar_url
	AuthorAvatarUrl *string `json:"author_avatar_url,omitempty"`

	// committer_avatar_url
	CommitterAvatarUrl *string `json:"committer_avatar_url,omitempty"`

	// only for DevCloud
	RelateUrl *[]RelatedCommitSimpleDto `json:"relate_url,omitempty"`

	// 标题
	NickName *string `json:"nick_name,omitempty"`

	// tenant_name
	TenantName *string `json:"tenant_name,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	Stats *CommitStatsDto `json:"stats,omitempty"`

	// **参数解释：** 状态。 **取值范围：** 不涉及。
	Status *string `json:"status,omitempty"`

	LastPipeline *PipelineBasicDto `json:"last_pipeline,omitempty"`

	// **参数解释：** 仓库ID。 **取值范围：** 不涉及。
	ProjectId *int32 `json:"project_id,omitempty"`

	MergeRequest *SimpleMergeRequestDetailDto `json:"merge_request,omitempty"`

	// **参数解释：** 代码变更内容。 **取值范围：** 不涉及。
	CodeChanges    *[]SimpleDiffDto `json:"code_changes,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o ShowCommitResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowCommitResponse struct{}"
	}

	return strings.Join([]string{"ShowCommitResponse", string(data)}, " ")
}
