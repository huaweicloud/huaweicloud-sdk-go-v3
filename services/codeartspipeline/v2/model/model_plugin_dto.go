package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PluginDto struct {

	// **参数解释**： 扩展插件唯一ID。可以通过[查询插件版本详情](ShowPluginVersion.xml)接口，获取响应参数中unique_id。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	UniqueId *string `json:"unique_id,omitempty"`

	// **参数解释**： 插件展示图标URL。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	IconUrl *string `json:"icon_url,omitempty"`

	// **参数解释**： 运行属性。 **约束限制**： 不涉及。 **取值范围**： - agent：基于agent运行。 - agentless：无需agent运行。 **默认取值**： 不涉及。
	RuntimeAttribution string `json:"runtime_attribution"`

	// **参数解释**： 插件名。 **约束限制**： 仅支持输入大小写英文字母、数字、'-'、'_'。 **取值范围**： 1到50位字符。 **默认取值**： 不涉及。
	PluginName string `json:"plugin_name"`

	// **参数解释**： 展示名。 **约束限制**： 仅支持输入大小写英文字母、中文、空格、数字、'-'、'_'、'.'。 **取值范围**： 1到50位字符。 **默认取值**： 不涉及。
	DisplayName string `json:"display_name"`

	// **参数解释**： 业务类型。 **约束限制**： 仅支持输入大小写英文字母、数字、'-'、'_'。 **取值范围**： 1到50位字符。 **默认取值**： 不涉及。
	BusinessType string `json:"business_type"`

	// **参数解释**： 插件业务类型展示名。 **约束限制**： 不涉及。 **取值范围**： - 构建。 - 代码检查。 - 部署。 - 测试。 - 通用。 **默认取值**： 不涉及。
	BusinessTypeDisplayName string `json:"business_type_display_name"`

	// **参数解释**： 插件描述。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description string `json:"description"`

	// **参数解释**： 是否私有插件。 **约束限制**： 不涉及。 **取值范围**： - 1：私有插件。 - 0：公开插件。 **默认取值**： 0。
	IsPrivate *int32 `json:"is_private,omitempty"`

	// **参数解释**： 局点。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Region *string `json:"region,omitempty"`

	// **参数解释**： 插件维护者。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Maintainers *string `json:"maintainers,omitempty"`

	// **参数解释**： 插件的组合类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PluginCompositionType *string `json:"plugin_composition_type,omitempty"`

	// **参数解释**： 用于区分新旧版数据版本。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ManifestVersion *string `json:"manifest_version,omitempty"`

	// **参数解释**： 插件版本号。 **约束限制**： 必须是类似 x.xx.xx（例如：1.0.2） 的格式，其中：x 是 1 到 2 位的数字（范围 0 到 99）。xx 是点后跟随的数字部分，且每部分可以是 1 位或 2 位数字。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Version string `json:"version"`

	// **参数解释**： 插件小版本版本号说明。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	VersionDescription *string `json:"version_description,omitempty"`

	ExecutionInfo *PluginDtoExecutionInfo `json:"execution_info"`

	// **参数解释**： 插件输出相关内容。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	OutputInfo *[]PluginDtoOutputInfo `json:"output_info,omitempty"`

	// **参数解释**： 输入信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InputInfo *[]PluginDtoInputInfo `json:"input_info,omitempty"`
}

func (o PluginDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginDto struct{}"
	}

	return strings.Join([]string{"PluginDto", string(data)}, " ")
}
