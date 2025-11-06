package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CommitMergeRequestDto 提交关联的合并请求详情
type CommitMergeRequestDto struct {

	// **参数解释：** 合并请求的ID。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 合并请求的序号。
	Iid *int32 `json:"iid,omitempty"`

	// **参数解释：** 合并请求的标题。
	Title *string `json:"title,omitempty"`

	// **参数解释：** 合并请求的详细描述。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 合并请求的状态。
	State *string `json:"state,omitempty"`

	// **参数解释：** 合并请求创建的时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 合并请求最后一次更新的时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	MergedBy *UserBasicDto `json:"merged_by,omitempty"`

	// **参数解释：** 合并请求被合并的时间。
	MergedAt *string `json:"merged_at,omitempty"`

	ClosedBy *UserBasicDto `json:"closed_by,omitempty"`

	// **参数解释：** 合并请求被关闭的时间。
	ClosedAt *string `json:"closed_at,omitempty"`

	// **参数解释：** 合并请求的目标分支名称。
	TargetBranch *string `json:"target_branch,omitempty"`

	// **参数解释：** 合并请求的源分支名称。
	SourceBranch *string `json:"source_branch,omitempty"`

	// **参数解释：** 合并请求中检视意见的数量。
	UserNotesCount *int32 `json:"user_notes_count,omitempty"`

	// **参数解释：** 合并请求的正向评分数量。
	Upvotes *int32 `json:"upvotes,omitempty"`

	// **参数解释：** 合并请求的负向评分数量。
	Downvotes *int32 `json:"downvotes,omitempty"`

	Author *UserBasicDto `json:"author,omitempty"`

	// **参数解释：** 合并请求的可合并人列表。
	Assignee *[]UserBasicDto `json:"assignee,omitempty"`

	// **参数解释：** 源仓库的唯一标识符。
	SourceRepositoryId *int32 `json:"source_repository_id,omitempty"`

	// **参数解释：** 目标仓库的唯一标识符。
	TargetRepositoryId *int32 `json:"target_repository_id,omitempty"`

	// **参数解释：** 合并请求关联的标签列表。
	Labels *[]string `json:"labels,omitempty"`

	// **参数解释：** 表示合并请求是否处于“工作进行中”状态。
	WorkInProgress *bool `json:"work_in_progress,omitempty"`

	Milestone *MilestoneBasicDto `json:"milestone,omitempty"`

	// **参数解释：** 表示是否在CI/CD管道成功时自动合并请求。
	MergeWhenPipelineSucceeds *bool `json:"merge_when_pipeline_succeeds,omitempty"`

	// **参数解释：** 合并请求的合并状态。
	MergeStatus *string `json:"merge_status,omitempty"`

	// **参数解释：** 合并请求的提交哈希值。
	Sha *string `json:"sha,omitempty"`

	// **参数解释：** 合并提交的哈希值。
	MergeCommitSha *string `json:"merge_commit_sha,omitempty"`

	// **参数解释：** 表示合并请求的讨论是否被锁定。
	DiscussionLocked *bool `json:"discussion_locked,omitempty"`

	// **参数解释：** 表示是否强制删除源分支。
	ForceRemoveSourceBranch *bool `json:"force_remove_source_branch,omitempty"`

	// **参数解释：** 表示是否应该删除源分支。
	ShouldRemoveSourceBranch *bool `json:"should_remove_source_branch,omitempty"`

	// **参数解释：** 表示是否允许协作者参与。
	AllowCollaboration *bool `json:"allow_collaboration,omitempty"`

	// **参数解释：** 表示是否允许维护者推送代码。
	AllowMaintainerToPush *bool `json:"allow_maintainer_to_push,omitempty"`

	// **参数解释：** 合并请求的网页URL链接。
	WebUrl *string `json:"web_url,omitempty"`

	TimeStats *IssuableTimeStatsDto `json:"time_stats,omitempty"`

	// **参数解释：** 表示是否在合并时将提交压缩为一个。
	Squash *bool `json:"squash,omitempty"`

	// **参数解释：** 合并请求的类型。
	MergeRequestType *string `json:"merge_request_type,omitempty"`
}

func (o CommitMergeRequestDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CommitMergeRequestDto struct{}"
	}

	return strings.Join([]string{"CommitMergeRequestDto", string(data)}, " ")
}
