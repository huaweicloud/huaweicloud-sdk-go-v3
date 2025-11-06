package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RelatedCommitVo struct {

	// 主键ID
	Id *string `json:"id,omitempty"`

	// 用户ID
	IamId *string `json:"iamId,omitempty"`

	// 用户名称
	UserName *string `json:"userName,omitempty"`

	// 仓库ID
	RepositoryId *string `json:"repositoryId,omitempty"`

	// 类型
	Type *string `json:"type,omitempty"`

	// 用户ID
	UserId *string `json:"userId,omitempty"`

	// 分支名称
	BranchName *string `json:"branchName,omitempty"`

	// Commit ID
	CommitId *string `json:"commitId,omitempty"`

	// Commit 短ID
	CommitShortId *string `json:"commitShortId,omitempty"`

	// 提交信息
	CommitMsg *string `json:"commitMsg,omitempty"`

	// 提交URL
	CommitUrl *string `json:"commitUrl,omitempty"`

	// 提交类型
	CommitType *string `json:"commitType,omitempty"`

	RelatedId *string `json:"relatedId,omitempty"`

	// 创建时间
	CreateAt *string `json:"createAt,omitempty"`

	// 更新时间
	UpdateAt *string `json:"updateAt,omitempty"`

	RelatedUrl *string `json:"relatedUrl,omitempty"`

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
