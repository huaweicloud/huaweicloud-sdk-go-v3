package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPipelineRunsPageBuildParams 构建参数
type ListPipelineRunsPageBuildParams struct {

	// 合并请求事件类型
	Action *string `json:"action,omitempty"`

	// 基于分支还是tag触发
	BuildType *string `json:"build_type,omitempty"`

	// 代码仓提交ID
	CommitId *string `json:"commit_id,omitempty"`

	// 运行事件类型
	EventType *string `json:"event_type,omitempty"`

	// 合并请求ID
	MergeId *string `json:"merge_id,omitempty"`

	// 代码仓提交信息
	Message *string `json:"message,omitempty"`

	// 源分支
	SourceBranch *string `json:"source_branch,omitempty"`

	// 标签
	Tag *string `json:"tag,omitempty"`

	// 目标分支
	TargetBranch *string `json:"target_branch,omitempty"`

	// codehub代码仓ID
	CodehubId *string `json:"codehub_id,omitempty"`

	// 代码仓https地址
	GitUrl *string `json:"git_url,omitempty"`

	// 源codehub代码仓ID
	SourceCodehubId *string `json:"source_codehub_id,omitempty"`

	// 源codehub代码仓地址
	SourceCodehubUrl *string `json:"source_codehub_url,omitempty"`

	// 源codehub代码仓http地址
	SourceCodehubHttpUrl *string `json:"source_codehub_http_url,omitempty"`
}

func (o ListPipelineRunsPageBuildParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelineRunsPageBuildParams struct{}"
	}

	return strings.Join([]string{"ListPipelineRunsPageBuildParams", string(data)}, " ")
}
