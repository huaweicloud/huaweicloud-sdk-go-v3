package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginBasicVo struct {

	// **参数解释**： 扩展插件名称。 **取值范围**： 1到50位字符。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 扩展插件名称。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 插件记录展示名称。 **取值范围**： 不涉及。
	DbRecordName *string `json:"db_record_name,omitempty"`

	// **参数解释**： 扩展插件版本号。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 扩展插件版本号说明。 **取值范围**： 不涉及。
	VersionDescription *string `json:"version_description,omitempty"`

	// **参数解释**： 扩展插件描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 扩展插件版本属性。 **取值范围**： - draft：草稿版本。 - formal：正式版本。
	VersionAttribution *string `json:"version_attribution,omitempty"`

	// **参数解释**： 扩展插件唯一ID。 **取值范围**： 不涉及。
	UniqueId *string `json:"unique_id,omitempty"`

	// **参数解释**： 扩展插件最后更新人。 **取值范围**： 不涉及。
	OpUser *string `json:"op_user,omitempty"`

	// **参数解释**： 扩展插件最后更新时间。 **取值范围**： 不涉及。
	OpTime *string `json:"op_time,omitempty"`

	// **参数解释**： 用于标识插件是否为多个step组成的组合插件。 **取值范围**： - multi：组合插件。 - single：非组合插件。
	PluginCompositionType *string `json:"plugin_composition_type,omitempty"`

	// **参数解释**： 扩展插件属性。 **取值范围**： - custom：自定义插件。 - official：官方插件。 - published：已发布的发布商插件。
	PluginAttribution *string `json:"plugin_attribution,omitempty"`

	// **参数解释**： 租户ID，用户的domainId。 **取值范围**： 32位字符，仅由数字和字母组成。
	WorkspaceId *string `json:"workspace_id,omitempty"`

	// **参数解释**： 扩展插件业务类型。 **取值范围**： 不涉及。
	BusinessType *string `json:"business_type,omitempty"`

	// **参数解释**： 扩展插件业务类型展示名称。 **取值范围**： 不涉及。
	BusinessTypeDisplayName *string `json:"business_type_display_name,omitempty"`

	// **参数解释**： 扩展插件维护人。 **取值范围**： 不涉及。
	Maintainers *string `json:"maintainers,omitempty"`

	// **参数解释**： 扩展插件图标地址。 **取值范围**： 不涉及。
	IconUrl *string `json:"icon_url,omitempty"`

	// **参数解释**： 扩展插件被流水线引用次数。 **取值范围**： 不涉及。
	ReferCount *int32 `json:"refer_count,omitempty"`

	// **参数解释**： 扩展插件被流水线使用次数。 **取值范围**： 不涉及。
	UsageCount *int32 `json:"usage_count,omitempty"`

	// **参数解释**： 运行属性。 **取值范围**： - agent：运行基于流水线agent。 - agentLess：运行无需流水线agent。
	RuntimeAttribution *string `json:"runtime_attribution,omitempty"`

	// **参数解释**： 扩展插件是否激活。 **取值范围**： - true：激活。 - false：未激活。
	Active *int32 `json:"active,omitempty"`

	// **参数解释**： 当前插件版本状态。 **取值范围**： 不涉及。
	VersionState *string `json:"version_state,omitempty"`

	// **参数解释**： 发布商ID。 **取值范围**： 不涉及。
	PublisherUniqueId *string `json:"publisher_unique_id,omitempty"`

	// **参数解释**： 创建者名称。 **取值范围**： 不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释**： 创建时间。 **取值范围**： 不涉及。
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**： 插件版本标识符。 **取值范围**： 不涉及。
	ManifestVersion *string `json:"manifest_version,omitempty"`
}

func (o PluginBasicVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginBasicVo struct{}"
	}

	return strings.Join([]string{"PluginBasicVo", string(data)}, " ")
}
