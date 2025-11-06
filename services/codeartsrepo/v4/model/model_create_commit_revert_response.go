package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateCommitRevertResponse Response Object
type CreateCommitRevertResponse struct {

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

	// 8位commitId
	ShortId *string `json:"short_id,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 标题
	Title *string `json:"title,omitempty"`

	// 作者头像url
	AuthorAvatarUrl *string `json:"author_avatar_url,omitempty"`

	// 提交人头像url
	CommitterAvatarUrl *string `json:"committer_avatar_url,omitempty"`

	// 以创建MR的形式Revert结果
	State *string `json:"state,omitempty"`

	// CherryPick临时分支名
	CherryPickBranchName *string `json:"cherry_pick_branch_name,omitempty"`

	// Revert临时分支名
	RevertBranchName *string `json:"revert_branch_name,omitempty"`

	// 变更请求 iid,仅变更请求返回
	Iid            *int32 `json:"iid,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o CreateCommitRevertResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommitRevertResponse struct{}"
	}

	return strings.Join([]string{"CreateCommitRevertResponse", string(data)}, " ")
}
