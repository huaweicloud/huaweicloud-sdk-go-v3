package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoOptionalSinglePluginVoData struct {

	// 名称
	Name *string `json:"name,omitempty"`

	// 展示名
	DisplayName *string `json:"display_name,omitempty"`

	// 属性
	PluginAttribution *string `json:"plugin_attribution,omitempty"`

	// 图标URL
	IconUrl *string `json:"icon_url,omitempty"`

	// 描述
	Description *string `json:"description,omitempty"`

	// 发布商ID
	PublisherId *string `json:"publisher_id,omitempty"`

	// 版本
	ManifestVersion *string `json:"manifest_version,omitempty"`
}

func (o PageInfoOptionalSinglePluginVoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoOptionalSinglePluginVoData struct{}"
	}

	return strings.Join([]string{"PageInfoOptionalSinglePluginVoData", string(data)}, " ")
}
