package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPipelineDetailResponse Response Object
type ShowPipelineDetailResponse struct {

	// 流水线ID
	Id *string `json:"id,omitempty"`

	// 流水线名称
	Name *string `json:"name,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 流水线版本
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// 局点
	Region *string `json:"region,omitempty"`

	// 所属租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 所属项目ID
	ProjectId *string `json:"project_id,omitempty"`

	// 所属微服务ID
	ComponentId *string `json:"component_id,omitempty"`

	// 是否为变更流水线
	IsPublish *bool `json:"is_publish,omitempty"`

	// 创建人ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 更新人ID
	UpdaterId *string `json:"updater_id,omitempty"`

	// 更新人名称
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 是否被当前用户收藏
	IsCollect *bool `json:"is_collect,omitempty"`

	// 流水线源
	Sources *[]PipelineSource `json:"sources,omitempty"`

	// 流水线自定义参数
	Variables *[]PipelineVariable `json:"variables,omitempty"`

	// 流水线定时任务设置
	Schedules *[]PipelineSchedule `json:"schedules,omitempty"`

	// 流水线事件触发设置
	Triggers *[]PipelineTrigger `json:"triggers,omitempty"`

	// 流水线所属分组ID
	GroupId *string `json:"group_id,omitempty"`

	// 流水线定义
	Definition     *string `json:"definition,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPipelineDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineDetailResponse", string(data)}, " ")
}
