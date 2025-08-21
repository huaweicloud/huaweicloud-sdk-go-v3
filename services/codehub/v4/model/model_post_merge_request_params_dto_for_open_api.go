package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PostMergeRequestParamsDtoForOpenApi struct {

	// 合并请求标题
	Title string `json:"title"`

	// 源分支
	SourceBranch string `json:"source_branch"`

	// 目标分支
	TargetBranch string `json:"target_branch"`

	// 目标仓库id
	TargetRepositoryId *int32 `json:"target_repository_id,omitempty"`

	// 打分模式评审人ids，使用逗号分隔
	ReviewerIds *string `json:"reviewer_ids,omitempty"`

	// 合并人ids，使用逗号分隔
	AssigneeIds *string `json:"assignee_ids,omitempty"`

	// 审核模式检视人ids，使用逗号分隔
	ApprovalReviewerIds *string `json:"approval_reviewer_ids,omitempty"`

	// 审核人ids，使用逗号分隔
	ApprovalApproversIds *string `json:"approval_approvers_ids,omitempty"`

	// 合并请求描述
	Description *string `json:"description,omitempty"`

	// 里程碑id
	MilestoneId *int32 `json:"milestone_id,omitempty"`

	// 标签名称，使用逗号分隔
	Labels *interface{} `json:"labels,omitempty"`

	// 合入后自动删除源分支
	ForceRemoveSourceBranch *bool `json:"force_remove_source_branch,omitempty"`

	// 压缩合并
	Squash *bool `json:"squash,omitempty"`

	// 压缩合并信息
	SquashCommitMessage *string `json:"squash_commit_message,omitempty"`

	// e2e单号
	WorkItemIds *[]string `json:"work_item_ids,omitempty"`

	// 使用临时分支创建合并请求
	IsUseTempBranch *bool `json:"is_use_temp_branch,omitempty"`

	// 只有合并人允许合入
	OnlyAssigneeCanMerge *bool `json:"only_assignee_can_merge,omitempty"`
}

func (o PostMergeRequestParamsDtoForOpenApi) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PostMergeRequestParamsDtoForOpenApi struct{}"
	}

	return strings.Join([]string{"PostMergeRequestParamsDtoForOpenApi", string(data)}, " ")
}
