package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CustomerPluginInfoResult struct {

	// **参数解释**: 插件名称。 **取值范围**: 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**: 是否安装此插件。 **取值范围**: 不涉及。
	Installed *bool `json:"installed,omitempty"`

	// **参数解释**: 插件端口。 **取值范围**: 不涉及。
	Port *string `json:"port,omitempty"`

	// **参数解释**: 插件版本。 **取值范围**: 不涉及。
	PluginVersion *string `json:"plugin_version,omitempty"`
}

func (o CustomerPluginInfoResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomerPluginInfoResult struct{}"
	}

	return strings.Join([]string{"CustomerPluginInfoResult", string(data)}, " ")
}
