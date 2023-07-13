package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PipelineTemplateSimpleVo struct {

	// 模板ID
	Id *string `json:"id,omitempty"`

	// 模板名称
	Name *string `json:"name,omitempty"`

	// 模板图标
	Icon *string `json:"icon,omitempty"`

	// 流水线结构定义版本，新版默认为3.0
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// 模板语言
	Language *string `json:"language,omitempty"`

	// 模板描述
	Description *string `json:"description,omitempty"`

	// 是否系统模板
	IsSystem *bool `json:"is_system,omitempty"`

	// 模板局点
	Region *string `json:"region,omitempty"`

	// 模板所属租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 模板创建人ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 模板创建人名称
	CreatorName *string `json:"creator_name,omitempty"`

	// 模板更新人ID
	UpdaterId *string `json:"updater_id,omitempty"`

	// 创建时间
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新时间
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 是否收藏
	IsCollect *bool `json:"is_collect,omitempty"`

	// 是否展示流水线源
	IsShowSource *string `json:"is_show_source,omitempty"`

	// 模板编排stages
	Stages *[]PipelineTemplateSimpleVoStages `json:"stages,omitempty"`
}

func (o PipelineTemplateSimpleVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PipelineTemplateSimpleVo struct{}"
	}

	return strings.Join([]string{"PipelineTemplateSimpleVo", string(data)}, " ")
}
