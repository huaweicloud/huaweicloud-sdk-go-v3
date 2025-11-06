package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBranchResponse Response Object
type CreateBranchResponse struct {

	// 分支名称
	Name *string `json:"name,omitempty"`

	// 是否为默认分支
	Default *bool `json:"default,omitempty"`

	// 用户是否为默认分支
	CanDelete *bool `json:"can_delete,omitempty"`

	// 是否为默认分支
	CanRead *bool `json:"can_read,omitempty"`

	// 是否为默认分支
	CanDownload *bool `json:"can_download,omitempty"`

	// 是否为默认分支
	CanPush *bool `json:"can_push,omitempty"`

	Commit *CommitDto `json:"commit,omitempty"`

	// 所有提交是否合入到默认分支
	Merged *bool `json:"merged,omitempty"`

	// 是否为保护分支
	Protected *bool `json:"protected,omitempty"`

	// 创建时间
	CreatedAt *string `json:"created_at,omitempty"`

	Creator *UserBasicDto `json:"creator,omitempty"`

	// 分支描述
	Description *string `json:"description,omitempty"`

	// 基于分支
	CreateSource *string `json:"create_source,omitempty"`

	// 基于分支是否存在
	CreateSourceExists *bool `json:"create_source_exists,omitempty"`

	LatestPipeline *PipelineBasicDto `json:"latest_pipeline,omitempty"`

	// 该分支正在开启的合并请求个数
	OpenedMrCount *int32 `json:"opened_mr_count,omitempty"`

	DivergingCommitCounts *DivergingCommitCounts `json:"diverging_commit_counts,omitempty"`
	HttpStatusCode        int                    `json:"-"`
}

func (o CreateBranchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBranchResponse struct{}"
	}

	return strings.Join([]string{"CreateBranchResponse", string(data)}, " ")
}
