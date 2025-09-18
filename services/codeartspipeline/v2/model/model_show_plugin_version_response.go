package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginVersionResponse Response Object
type ShowPluginVersionResponse struct {

	// **参数解释**： 扩展插件名称。 **取值范围**： 1到50位字符。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 扩展插件展示名称。 **取值范围**： 不涉及。
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 扩展插件最后更新人。 **取值范围**： 不涉及。
	OpUser *string `json:"op_user,omitempty"`

	// **参数解释**： 扩展插件最后更新时间。 **取值范围**： 不涉及。
	OpTime *string `json:"op_time,omitempty"`

	// **参数解释**： 扩展插件版本号。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 扩展插件唯一ID。 **取值范围**： 32位字符，由数字和字母组成。
	UniqueId *string `json:"unique_id,omitempty"`

	// **参数解释**： 扩展插件版本说明。 **取值范围**： 32位字符，由数字和字母组成。
	VersionDescription *string `json:"version_description,omitempty"`

	// **参数解释**： 扩展插件版本属性。 **取值范围**： - draft：草稿版本。 - formal：正式版本。
	VersionAttribution *string `json:"version_attribution,omitempty"`

	// **参数解释**： 用于标识扩展插件是否为多个step组成的组合插件。 **取值范围**： - multi：组合插件。 - single：非组合插件。
	PluginCompositionType *string `json:"plugin_composition_type,omitempty"`

	// **参数解释**： 扩展插件属性。 **取值范围**： - custom：自定义插件。 - official：官方插件。 - published：已发布的发布商插件。
	PluginAttribution *string `json:"plugin_attribution,omitempty"`

	// **参数解释**： 插件输入项详细信息。 **取值范围**： 不涉及。
	InputInfo *[]PluginInstanceVoInputInfo `json:"input_info,omitempty"`

	PluginExecution *PluginExecutionVo `json:"plugin_execution,omitempty"`

	// **参数解释**： 运行属性。 **取值范围**： - agent：运行基于流水线agent。 - agentLess：运行无需流水线agent。
	RuntimeAttribution *string `json:"runtime_attribution,omitempty"`
	HttpStatusCode     int     `json:"-"`
}

func (o ShowPluginVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowPluginVersionResponse", string(data)}, " ")
}
