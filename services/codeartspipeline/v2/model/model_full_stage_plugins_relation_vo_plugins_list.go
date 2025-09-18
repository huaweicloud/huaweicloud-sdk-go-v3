package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type FullStagePluginsRelationVoPluginsList struct {

	// **参数解释**： 唯一ID。 **取值范围**： 不涉及。
	UniqueId *string `json:"unique_id,omitempty"`

	// **参数解释**： 展示名。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 插件名。 **取值范围**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 当前插件对后续选择使用的流水线是否为禁用状态，默认为false。 **取值范围**： - true：被禁用。 - false：未被禁用。
	Disabled *bool `json:"disabled,omitempty"`

	// **参数解释**： 插件记录展示名称。 **取值范围**： 不涉及。
	DbRecordName *string `json:"db_record_name,omitempty"`

	// **参数解释**： 流水线stage下的分组名称。 **取值范围**： 不涉及。
	GroupName *string `json:"group_name,omitempty"`

	// **参数解释**： 流水线stage下的分组类型。 **取值范围**： 不涉及。
	GroupType *string `json:"group_type,omitempty"`

	// **参数解释**： 标识是否为官方插件。 **取值范围**： 不涉及。
	PluginAttribution *string `json:"plugin_attribution,omitempty"`

	// **参数解释**： 标识是否为多个step组成的组。 **取值范围**： - single：单step插件。 - multi：组合插件。
	PluginCompositionType *string `json:"plugin_composition_type,omitempty"`

	// **参数解释**： 运行属性。 **取值范围**： - agent：运行基于流水线agent。 - agentLess：运行无需流水线agent。
	RuntimeAttribution *string `json:"runtime_attribution,omitempty"`

	// **参数解释**： 基础插件列表。 **取值范围**： 不涉及。
	AllSteps *[]FullStagePluginsRelationVoAllSteps `json:"all_steps,omitempty"`

	// **参数解释**： 扩展插件描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 标识是否为一个草稿。 **取值范围**： 不涉及。
	VersionAttribution *string `json:"version_attribution,omitempty"`

	// **参数解释**： 扩展插件图标URL。 **取值范围**： 不涉及。
	IconUrl *string `json:"icon_url,omitempty"`

	// **参数解释**： 标识是否可继续进行添加步骤，默认是1，可进行添加。 **取值范围**： - 0：不可继续进行添加步骤。 - 1：可继续进行添加步骤。
	MultiStepEditable *int32 `json:"multi_step_editable,omitempty"`

	// **参数解释**： 使用位置。 **取值范围**： 不涉及。
	Location *string `json:"location,omitempty"`

	// **参数解释**： 发布商ID。 **取值范围**： 不涉及。
	PublisherUniqueId *string `json:"publisher_unique_id,omitempty"`

	// **参数解释**： 插件版本标识符。 **取值范围**： 不涉及。
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// **参数解释**： 标识是否为标准化的插件。 **取值范围**： - true：是标准化的插件。 - false：不是标准化的插件。
	Standard *bool `json:"standard,omitempty"`
}

func (o FullStagePluginsRelationVoPluginsList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FullStagePluginsRelationVoPluginsList struct{}"
	}

	return strings.Join([]string{"FullStagePluginsRelationVoPluginsList", string(data)}, " ")
}
