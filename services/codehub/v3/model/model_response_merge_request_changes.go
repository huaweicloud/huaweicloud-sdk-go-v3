package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResponseMergeRequestChanges struct {

	// 合并请求id
	Id *int32 `json:"id,omitempty"`

	// 合并请求iid
	Iid *int32 `json:"iid,omitempty"`

	// 仓库id
	ProjectId *int32 `json:"project_id,omitempty"`

	// 合并请求标题
	Title *string `json:"title,omitempty"`

	// 合并请求描述
	Description *string `json:"description,omitempty"`

	// 合并请求状态
	State *string `json:"state,omitempty"`

	// 合并请求创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	// 合并请求更新时间
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 合并请求目标分支
	TargetBranch *string `json:"target_branch,omitempty"`

	// 合并请求源分支
	SourceBranch *string `json:"source_branch,omitempty"`

	// 检视意见数量
	UserNotesCount *int32 `json:"user_notes_count,omitempty"`

	// upvotes
	Upvotes *int32 `json:"upvotes,omitempty"`

	// downvotes
	Downvotes *int32 `json:"downvotes,omitempty"`

	Author *Author `json:"author,omitempty"`

	// 合并请求源仓库id
	SourceProjectId *int32 `json:"source_project_id,omitempty"`

	// 合并请求目标仓库id
	TargetProjectId *int32 `json:"target_project_id,omitempty"`

	// 指定仓库的标签列表
	Labels *[]string `json:"labels,omitempty"`

	// 合并请求是否为wip状态
	WorkInProgress *bool `json:"work_in_progress,omitempty"`

	// 合并请求是否开启流水线成功后自动合入
	MergeWhenPipelineSucceeds *bool `json:"merge_when_pipeline_succeeds,omitempty"`

	// 合并请求合入状态
	MergeStatus *string `json:"merge_status,omitempty"`

	// 合并请求的head sha
	Sha *string `json:"sha,omitempty"`

	// 合并请求合入后是否应该移除源分支
	ShouldRemoveSourceBranch *bool `json:"should_remove_source_branch,omitempty"`

	// 合并请求合入后是否移除源分支
	ForceRemoveSourceBranch *bool `json:"force_remove_source_branch,omitempty"`

	// 合并请求url
	WebUrl *string `json:"web_url,omitempty"`

	// 合并请求是否开启squash合并
	Squash *bool `json:"squash,omitempty"`

	// 合并请求类型
	MergeRequestType *string `json:"merge_request_type,omitempty"`

	// 是否订阅
	Subscribed *bool `json:"subscribed,omitempty"`

	// 合并请求变更文件数量
	ChangesCount *string `json:"changes_count,omitempty"`

	// 合并请求最新构建开始时间
	LatestBuildStartedAt *string `json:"latest_build_started_at,omitempty"`

	// 合并请求最新构建结束时间
	LatestBuildFinishedAt *string `json:"latest_build_finished_at,omitempty"`

	// first deployed to production at
	FirstDeployedToProductionAt *string `json:"first_deployed_to_production_at,omitempty"`

	Pipeline *PipelineBasicDto `json:"pipeline,omitempty"`

	DiffRefs *DiffRefsDto `json:"diff_refs,omitempty"`

	// 合并请求操作错误信息
	MergeError *string `json:"merge_error,omitempty"`

	// 合并请求是否rebase中
	RebaseInProgress *bool `json:"rebase_in_progress,omitempty"`

	// 合并请求落后提交数量
	DivergedCommitsCount *int32 `json:"diverged_commits_count,omitempty"`

	User *MergeRequestListDtoUser `json:"user,omitempty"`

	// 合并请求增加行数
	AddedLines *int32 `json:"added_lines,omitempty"`

	// 合并请求删除行数
	RemovedLines *int32 `json:"removed_lines,omitempty"`

	// 合并请求文件变更
	Changes *[]MergeRequestDiffFileDto `json:"changes,omitempty"`

	SourceProject *ProjectSimpleDto `json:"source_project,omitempty"`
}

func (o ResponseMergeRequestChanges) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResponseMergeRequestChanges struct{}"
	}

	return strings.Join([]string{"ResponseMergeRequestChanges", string(data)}, " ")
}
