package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MergeRequestCommitDto 合并请求commit列表详情
type MergeRequestCommitDto struct {

	// commit id
	Id *string `json:"id,omitempty"`

	// commit 短id
	ShortId *string `json:"short_id,omitempty"`

	// 提交标题
	Title *string `json:"title,omitempty"`

	// 提交信息
	Message *string `json:"message,omitempty"`

	// commit 作者名称
	AuthorName *string `json:"author_name,omitempty"`

	// 用户名
	Name *string `json:"name,omitempty"`

	// 用户名
	UserName *string `json:"user_name,omitempty"`

	// 租户名
	TenantName *string `json:"tenant_name,omitempty"`

	// 昵称
	NickName *string `json:"nick_name,omitempty"`

	// 最初commit 提交日期（本地提交日期）
	AuthoredDate *string `json:"authored_date,omitempty"`

	// commit提交日期（推送至仓库日期）
	CommittedDate *string `json:"committed_date,omitempty"`

	// commit 提交者名称
	CommitterName *string `json:"committer_name,omitempty"`

	// pgp key id
	GpgPrimaryKeyId *string `json:"gpg_primary_key_id,omitempty"`

	// gpg公钥验证是否开启
	OpenGpgVerified *bool `json:"open_gpg_verified,omitempty"`

	// gpg公钥验证是否通过
	VerificationStatus *bool `json:"verification_status,omitempty"`

	// 提交父commit节点
	ParentIds *[]string `json:"parent_ids,omitempty"`

	// commit 数据库记录创建时间
	CreatedAt *string `json:"created_at,omitempty"`
}

func (o MergeRequestCommitDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestCommitDto struct{}"
	}

	return strings.Join([]string{"MergeRequestCommitDto", string(data)}, " ")
}
