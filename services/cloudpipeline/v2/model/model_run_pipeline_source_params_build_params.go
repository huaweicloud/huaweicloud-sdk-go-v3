package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RunPipelineSourceParamsBuildParams 构建参数
type RunPipelineSourceParamsBuildParams struct {

	// 合并请求事件类型
	Action *string `json:"action,omitempty"`

	// 基于分支还是tag运行
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
}

func (o RunPipelineSourceParamsBuildParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RunPipelineSourceParamsBuildParams struct{}"
	}

	return strings.Join([]string{"RunPipelineSourceParamsBuildParams", string(data)}, " ")
}
