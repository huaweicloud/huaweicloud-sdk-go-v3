package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelatedCommitVo struct {

	// 主键ID
	Id *string `json:"id,omitempty"`

	// 用户ID
	IamId *string `json:"iam_id,omitempty"`

	// 用户名称
	UserName *string `json:"user_name,omitempty"`

	// 仓库ID
	RepositoryId *string `json:"repository_id,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 用户ID
	UserId *string `json:"user_id,omitempty"`

	// 分支名称
	BranchName *string `json:"branch_name,omitempty"`

	// Commit ID
	CommitId *string `json:"commit_id,omitempty"`

	// Commit 短ID
	CommitShortId *string `json:"commit_shortId,omitempty"`

	// 提交信息
	CommitMsg *string `json:"commit_msg,omitempty"`

	// 提交URL
	CommitUrl *string `json:"commit_url,omitempty"`

	// 提交类型
	CommitType *string `json:"commit_type,omitempty"`

	RelatedId *string `json:"related_id,omitempty"`

	// 创建时间
	CreateAt *string `json:"create_at,omitempty"`

	// 更新时间
	UpdateAt *string `json:"update_at,omitempty"`

	RelatedUrl *string `json:"related_url,omitempty"`

	// 描述
	Message *string `json:"message,omitempty"`
}

func (o RelatedCommitVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RelatedCommitVo struct{}"
	}

	return strings.Join([]string{"RelatedCommitVo", string(data)}, " ")
}
