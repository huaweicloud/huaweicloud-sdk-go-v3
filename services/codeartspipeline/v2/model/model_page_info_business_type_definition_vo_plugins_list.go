package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoBusinessTypeDefinitionVoPluginsList struct {

	// **参数解释**： 唯一ID。 **取值范围**： 不涉及。
	UniqueId *string `json:"unique_id,omitempty"`

	// **参数解释**： 展示名。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 插件名。 **取值范围**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 是否禁用。 **取值范围**： - true：禁用。 - false：未禁用。
	Disabled *bool `json:"disabled,omitempty"`

	// **参数解释**： 组名。 **取值范围**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 组类型。 **取值范围**： 不涉及。
	GroupType *string `json:"group_type,omitempty"`

	// **参数解释**： 是否标准化的插件。 **取值范围**： - true：是。 - false：否。
	Standard *bool `json:"standard,omitempty"`

	// **参数解释**： 插件记录名称。 **取值范围**： 不涉及。
	DbRecordName *string `json:"db_record_name,omitempty"`

	// **参数解释**： 插件属性。 **取值范围**： 不涉及。
	PluginAttribution *string `json:"plugin_attribution,omitempty"`

	// **参数解释**： 组合插件。 **取值范围**： 不涉及。
	PluginCompositionType *string `json:"plugin_composition_type,omitempty"`

	// **参数解释**： 插件运行属性。 **取值范围**： 不涉及。
	RuntimeAttribution *string `json:"runtime_attribution,omitempty"`

	// **参数解释**： 基础插件列表。 **取值范围**： 不涉及。
	AllSteps *[]PageInfoBusinessTypeDefinitionVoAllSteps `json:"all_steps,omitempty"`

	// **参数解释**： 插件描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 插件版本属性。 **取值范围**： 不涉及。
	VersionAttribution *string `json:"version_attribution,omitempty"`

	// **参数解释**： 插件图标URL。 **取值范围**： 不涉及。
	IconUrl *string `json:"icon_url,omitempty"`

	// **参数解释**： 插件可编辑。 **取值范围**： 不涉及。
	MultiStepEditable *int32 `json:"multi_step_editable,omitempty"`

	// **参数解释**： 插件地址。 **取值范围**： 不涉及。
	Location *string `json:"location,omitempty"`

	// **参数解释**： 插件发布商ID。 **取值范围**： 不涉及。
	PublisherUniqueId *string `json:"publisher_unique_id,omitempty"`

	// **参数解释**： 插件版本。 **取值范围**： 不涉及。
	ManifestVersion *string `json:"manifest_version,omitempty"`
}

func (o PageInfoBusinessTypeDefinitionVoPluginsList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoBusinessTypeDefinitionVoPluginsList struct{}"
	}

	return strings.Join([]string{"PageInfoBusinessTypeDefinitionVoPluginsList", string(data)}, " ")
}
