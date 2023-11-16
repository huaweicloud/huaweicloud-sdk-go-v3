package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListPipelinesPagePipelines struct {

	// 流水线ID
	PipelineId *string `json:"pipeline_id,omitempty"`

	// 流水线名称
	Name *string `json:"name,omitempty"`

	// 项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 项目名称
	ProjectName *string `json:"project_name,omitempty"`

	// 组件ID
	ComponentId *string `json:"component_id,omitempty"`

	// 是否为变更流水线
	IsPublish *bool `json:"is_publish,omitempty"`

	// 是否收藏此流水线
	IsCollect *bool `json:"is_collect,omitempty"`

	// 流水线版本
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	LatestRun *ListPipelinesPageLatestRun `json:"latest_run,omitempty"`

	// 旧版转新版标识
	ConvertSign *int32 `json:"convert_sign,omitempty"`
}

func (o ListPipelinesPagePipelines) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPipelinesPagePipelines struct{}"
	}

	return strings.Join([]string{"ListPipelinesPagePipelines", string(data)}, " ")
}
