package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPipelineTemplateDetailResponse struct {

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

	// 所属局点
	Region *string `json:"region,omitempty"`

	// 所属租户ID
	DomainId *string `json:"domain_id,omitempty"`

	// 使用的自定义参数
	Variables *[]CustomVariable `json:"variables,omitempty"`

	// 创建人ID
	CreatorId *string `json:"creator_id,omitempty"`

	// 编辑人ID
	UpdaterId *string `json:"updater_id,omitempty"`

	// 创建日期
	CreateTime *int64 `json:"create_time,omitempty"`

	// 更新日期
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 是否收藏
	IsCollect *string `json:"is_collect,omitempty"`

	// 是否显示流水线源
	IsShowSource *bool `json:"is_show_source,omitempty"`

	// 模板编排json，包含stages
	Definition     *string `json:"definition,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPipelineTemplateDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPipelineTemplateDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowPipelineTemplateDetailResponse", string(data)}, " ")
}
