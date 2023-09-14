package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PipelineLatestRunArtifactParams 运行制品源参数
type PipelineLatestRunArtifactParams struct {

	// 包版本
	Version *string `json:"version,omitempty"`

	// 过滤分支
	BranchFilter *string `json:"branch_filter,omitempty"`

	// 包名称
	PackageName *string `json:"package_name,omitempty"`

	// docker组织信息
	Organization *string `json:"organization,omitempty"`
}

func (o PipelineLatestRunArtifactParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineLatestRunArtifactParams struct{}"
	}

	return strings.Join([]string{"PipelineLatestRunArtifactParams", string(data)}, " ")
}
