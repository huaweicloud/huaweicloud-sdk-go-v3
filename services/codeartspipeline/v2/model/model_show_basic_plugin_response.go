package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBasicPluginResponse Response Object
type ShowBasicPluginResponse struct {

	// **参数解释**： 插件类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 插件名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 插件展示名。 **取值范围**： 不涉及。
	FriendlyName *string `json:"friendly_name,omitempty"`

	// **参数解释**： 业务类型。 **取值范围**： 不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释**： 插件描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 插件版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 插件版本说明。 **取值范围**： 不涉及。
	VersionDescription *string `json:"version_description,omitempty"`

	// **参数解释**： 输入信息。 **取值范围**： 不涉及。
	Inputs *[]NewExtensionInputs `json:"inputs,omitempty"`

	// **参数解释**： 数据源绑定信息。 **取值范围**： 不涉及。
	DataSourceBindings *[]NewExtensionDataSourceBindings `json:"data_source_bindings,omitempty"`

	// **参数解释**： 输出信息。 **取值范围**： 不涉及。
	Outputs *[]NewExtensionOutputs `json:"outputs,omitempty"`

	Execution      *NewExtensionExecution `json:"execution,omitempty"`
	HttpStatusCode int                    `json:"-"`
}

func (o ShowBasicPluginResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBasicPluginResponse struct{}"
	}

	return strings.Join([]string{"ShowBasicPluginResponse", string(data)}, " ")
}
