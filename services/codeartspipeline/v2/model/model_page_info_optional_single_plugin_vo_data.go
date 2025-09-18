package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoOptionalSinglePluginVoData struct {

	// **参数解释**： 扩展插件唯一标识。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 扩展插件在插件市场和流水线显示的名称。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 标识是否为官方插件。 **取值范围**： 不涉及。
	PluginAttribution *string `json:"plugin_attribution,omitempty"`

	// **参数解释**： 图标URL。 **取值范围**： 不涉及。
	IconUrl *string `json:"icon_url,omitempty"`

	// **参数解释**： 插件描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 发布商ID。 **取值范围**： 不涉及。
	PublisherId *string `json:"publisher_id,omitempty"`

	// **参数解释**： 版本。 **取值范围**： 不涉及。
	ManifestVersion *string `json:"manifest_version,omitempty"`
}

func (o PageInfoOptionalSinglePluginVoData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoOptionalSinglePluginVoData struct{}"
	}

	return strings.Join([]string{"PageInfoOptionalSinglePluginVoData", string(data)}, " ")
}
