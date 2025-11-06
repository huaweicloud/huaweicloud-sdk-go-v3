package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MergeRequestListBasicDto 合并请求基本信息
type MergeRequestListBasicDto struct {

	// **参数解释：** 合并请求ID。
	Id *int32 `json:"id,omitempty"`

	// **参数解释：** 合并请求位于当前仓库序号。
	Iid *int32 `json:"iid,omitempty"`

	// **参数解释：** 合并请求标题。
	Title *string `json:"title,omitempty"`

	// **参数解释：** 合并请求源分支。
	SourceBranch *string `json:"source_branch,omitempty"`

	// **参数解释：** 合并请求目标分支。
	TargetBranch *string `json:"target_branch,omitempty"`

	// **参数解释：** 合并请求状态。
	State *string `json:"state,omitempty"`

	// **参数解释：** 合并请求创建时间。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释：** 合并请求更新时间。
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释：** 合并请求源仓库ID。
	SourceRepositoryId *int32 `json:"source_repository_id,omitempty"`

	// **参数解释：** 合并请求检视模式。
	ReviewMode *string `json:"review_mode,omitempty"`

	Author *UserBasicDto `json:"author,omitempty"`

	// **参数解释：** 合并请求关闭时间。
	ClosedAt *string `json:"closed_at,omitempty"`

	ClosedBy *UserBasicDto `json:"closed_by,omitempty"`

	// **参数解释：** 合并请求合并时间。
	MergedAt *string `json:"merged_at,omitempty"`

	MergedBy *UserBasicDto `json:"merged_by,omitempty"`

	// **参数解释：** 合并请求流水线状态。
	PipelineStatus *string `json:"pipeline_status,omitempty"`

	// **参数解释：** 合并请求代码质量状态。
	CodequalityStatus *string `json:"codequality_status,omitempty"`

	// **参数解释：** 合并请求流水线状态。
	PipelineStatusWithCodeQuality *string `json:"pipeline_status_with_code_quality,omitempty"`

	// **参数解释：** 合并请求检视意见。
	Notes *int32 `json:"notes,omitempty"`

	SourceRepository *ProjectSimpleDto `json:"source_repository,omitempty"`

	TargetRepository *ProjectSimpleDto `json:"target_repository,omitempty"`

	// **参数解释：** 合并请求URL地址。
	WebUrl *string `json:"web_url,omitempty"`

	// **参数解释：** 合并请求新增代码行数。
	AddedLines *int32 `json:"added_lines,omitempty"`

	// **参数解释：** 合并请求删除代码行数。
	RemovedLines *int32 `json:"removed_lines,omitempty"`

	// **参数解释：** 合并请求检视模式。
	MergeRequestType *string `json:"merge_request_type,omitempty"`

	// **参数解释：** 合并请求git地址。
	SourceGitUrl *string `json:"source_git_url,omitempty"`

	// **参数解释：** 合并请求标签。
	Labels *[]map[string]interface{} `json:"labels,omitempty"`

	// **参数解释：** 合并请求分数。
	Score *int32 `json:"score,omitempty"`

	// **参数解释：** 合并请求最小合入分数。
	MinMergedScore *int32 `json:"min_merged_score,omitempty"`

	// **参数解释：** 合并请求源项目ID。
	SourceProductId *string `json:"source_product_id,omitempty"`

	// **参数解释：** 合并请求目标项目ID。
	TargetProductId *string `json:"target_product_id,omitempty"`

	// **参数解释：** 合并请求项目名。
	ProductName *string `json:"product_name,omitempty"`

	NotesCount *NotesCountDto `json:"notes_count,omitempty"`

	// **参数解释：** 合并请求审核结果。
	ModerationResult *bool `json:"moderation_result,omitempty"`

	// **参数解释：** 合并请求审核时间。
	ModerationTime *int64 `json:"moderation_time,omitempty"`

	// **参数解释：** 合并请求审核状态。
	ModerationStatus *int32 `json:"moderation_status,omitempty"`
}

func (o MergeRequestListBasicDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MergeRequestListBasicDto struct{}"
	}

	return strings.Join([]string{"MergeRequestListBasicDto", string(data)}, " ")
}
