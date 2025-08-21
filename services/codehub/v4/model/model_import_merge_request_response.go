package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ImportMergeRequestResponse Response Object
type ImportMergeRequestResponse struct {

	// 合并请求id
	Id *int32 `json:"id,omitempty"`

	// 合并请求iid
	Iid *int32 `json:"iid,omitempty"`

	// 目标仓库id
	RepositoryId *int32 `json:"repository_id,omitempty"`

	// 合并请求标题
	Title *string `json:"title,omitempty"`

	// 合并请求描述
	Description *string `json:"description,omitempty"`

	// 合并请求状态
	State *string `json:"state,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 源分支
	SourceBranch *string `json:"source_branch,omitempty"`

	// 目标分支
	TargetBranch *string `json:"target_branch,omitempty"`

	// 源分支是否为保护分支
	IsSourceBranchProtected *bool `json:"is_source_branch_protected,omitempty"`

	// 源分支
	DevcloudSourceBranch *string `json:"devcloud_source_branch,omitempty"`

	Author *UserBasicExternalDto `json:"author,omitempty"`

	// 源仓库id
	SourceRepositoryId *int32 `json:"source_repository_id,omitempty"`

	// 目标仓库id
	TargetRepositoryId *int32 `json:"target_repository_id,omitempty"`

	// 源项目id
	SourceProjectId *string `json:"source_project_id,omitempty"`

	// 目标项目id
	TargetProjectId *string `json:"target_project_id,omitempty"`

	// 标签
	Labels *[]interface{} `json:"labels,omitempty"`

	// WIP状态
	WorkInProgress *bool `json:"work_in_progress,omitempty"`

	Milestone *MilestoneBasicDto `json:"milestone,omitempty"`

	// 流水线成功后自动合入
	MergeWhenBuildSucceeds *bool `json:"merge_when_build_succeeds,omitempty"`

	// 可合并状态
	MergeStatus *string `json:"merge_status,omitempty"`

	// 当前合并请求最新commit
	Sha *string `json:"sha,omitempty"`

	// 合入commit节点
	MergeCommitSha *string `json:"merge_commit_sha,omitempty"`

	// 订阅
	Subscribed *bool `json:"subscribed,omitempty"`

	MergedBy *UserBasicExternalDto `json:"merged_by,omitempty"`

	// 合并时间
	MergedAt *string `json:"merged_at,omitempty"`

	ClosedBy *UserBasicExternalDto `json:"closed_by,omitempty"`

	// 关闭时间
	ClosedAt *string `json:"closed_at,omitempty"`

	// 检视意见数量
	UserNotesCount *int32 `json:"user_notes_count,omitempty"`

	// 合入后删除源分支
	ForceRemoveSourceBranch *bool `json:"force_remove_source_branch,omitempty"`

	// 主页url
	WebUrl *string `json:"web_url,omitempty"`

	MergeRequestDiff *MergeRequestDiffExternalDto `json:"merge_request_diff,omitempty"`

	// 评审人数量
	MergeRequestReviewersCount *int32 `json:"merge_request_reviewers_count,omitempty"`

	// 打分
	MergeRequestReviewCount *int32 `json:"merge_request_review_count,omitempty"`

	// 评审人列表
	MergeRequestReviewerList *[]MergeRequestReviewerExternalDto `json:"merge_request_reviewer_list,omitempty"`

	// 合并人列表
	MergeRequestAssigneeList *[]UserBasicExternalDto `json:"merge_request_assignee_list,omitempty"`

	// 检视意见
	Notes *int32 `json:"notes,omitempty"`

	// 代码检查状态
	CodecheckState *int32 `json:"codecheck_state,omitempty"`

	// 代码检查问题数
	CodecheckDefectCount *int32 `json:"codecheck_defect_count,omitempty"`

	// 合并请求关联单数量
	MergeRequestRelatedWorkItems *[]MergeRequestRelatedWorkItemDto `json:"merge_request_related_work_items,omitempty"`

	// 源分支落后commit数
	DivergedCommitsCount *int32 `json:"diverged_commits_count,omitempty"`

	// 送审结果
	ModerationResult *bool `json:"moderation_result,omitempty"`

	// 送审时间时间戳
	ModerationTime *int64 `json:"moderation_time,omitempty"`

	// 送审状态码
	ModerationStatus *int32 `json:"moderation_status,omitempty"`

	// 是否使用临时分支
	IsUseTempBranch *bool `json:"is_use_temp_branch,omitempty"`

	// 检视模式
	ReviewMode *string `json:"review_mode,omitempty"`

	// squash合入
	Squash *bool `json:"squash,omitempty"`

	// 审核模式审核人
	ApprovalMergeRequestApprovers *[]ApprovalUserDto `json:"approval_merge_request_approvers,omitempty"`

	// 审核模式检视人
	ApprovalMergeRequestReviewers *[]ApprovalUserDto `json:"approval_merge_request_reviewers,omitempty"`

	SourceRepository *ProjectSimpleDto `json:"source_repository,omitempty"`

	TargetRepository *ProjectSimpleDto `json:"target_repository,omitempty"`

	// 源分支存在
	IsSourceBranchExist *bool `json:"is_source_branch_exist,omitempty"`

	// 合并请求类型
	MergeRequestType *string `json:"merge_request_type,omitempty"`

	// squash提交信息
	SquashCommitMessage *string `json:"squash_commit_message,omitempty"`

	// 自动合入
	AutoMerge *bool `json:"auto_merge,omitempty"`

	// MR原始错误信息
	MergeError *string `json:"merge_error,omitempty"`

	JsonMergeError *MergeErrorDto `json:"json_merge_error,omitempty"`

	// 是否正在rebase
	RebaseInProgress *bool `json:"rebase_in_progress,omitempty"`
	HttpStatusCode   int   `json:"-"`
}

func (o ImportMergeRequestResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ImportMergeRequestResponse struct{}"
	}

	return strings.Join([]string{"ImportMergeRequestResponse", string(data)}, " ")
}
