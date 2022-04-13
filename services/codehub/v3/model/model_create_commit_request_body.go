package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateCommitRequestBody struct {
	// 目标分支

	Branch string `json:"branch"`
	// 提交信息

	CommitMessage string `json:"commit_message"`
	// 创建分支时，新的分支名

	StartBranch *string `json:"start_branch,omitempty"`
	// 提交处理列表

	Actions []CommitAction `json:"actions"`
	// 提交作者的电子邮件地址

	AuthorEmail *string `json:"author_email,omitempty"`
	// 提交作者的名称

	AuthorName *string `json:"author_name,omitempty"`
	// 是否包括提交统计信息。默认值为true

	Stats *bool `json:"stats,omitempty"`
	// 是否覆盖目标分支。当true时，使用基于start_branch的新提交覆盖目标分支

	Force *string `json:"force,omitempty"`
}

func (o CreateCommitRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateCommitRequestBody struct{}"
	}

	return strings.Join([]string{"CreateCommitRequestBody", string(data)}, " ")
}
