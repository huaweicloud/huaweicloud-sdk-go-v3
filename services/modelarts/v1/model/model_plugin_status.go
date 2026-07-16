package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PluginStatus 插件状态信息。
type PluginStatus struct {

	// **参数解释**： 插件实例的状态。 **取值范围**：可选值如下： - Pending：安装中，表示插件正在安装中。 - Running：运行中，表示插件全部实例状态都在运行中，插件正常使用。 - Updating：升级中，表示插件正在更新中。 - Abnormal：不可用，表示插件状态异常，插件不可使用。可单击状态查看失败原因。 - Deleting：删除中，表示插件正在删除中。
	Phase *string `json:"phase,omitempty"`

	// **参数解释**： 插件实例的版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 插件实例安装失败的详细信息。 **取值范围**： 不涉及。
	Reason *string `json:"reason,omitempty"`

	// **参数解释**： 插件实例的安装参数（各插件不同）。 **取值范围**： 不涉及。
	Values *string `json:"values,omitempty"`

	// **参数解释**： 插件实例占用的资源量。
	Resources *[]PluginResources `json:"resources,omitempty"`
}

func (o PluginStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PluginStatus struct{}"
	}

	return strings.Join([]string{"PluginStatus", string(data)}, " ")
}
