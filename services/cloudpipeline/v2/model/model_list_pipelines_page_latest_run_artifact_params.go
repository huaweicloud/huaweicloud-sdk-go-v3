package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 制品源参数
type ListPipelinesPageLatestRunArtifactParams struct {

	// 包版本
	Version *string `json:"version,omitempty"`

	// 过滤分支
	BranchFilter *string `json:"branch_filter,omitempty"`

	// 包名称
	PackageName *string `json:"package_name,omitempty"`

	// docker组织
	Organization *string `json:"organization,omitempty"`
}

func (o ListPipelinesPageLatestRunArtifactParams) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesPageLatestRunArtifactParams struct{}"
	}

	return strings.Join([]string{"ListPipelinesPageLatestRunArtifactParams", string(data)}, " ")
}
