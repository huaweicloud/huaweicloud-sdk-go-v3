package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type InstallPluginRequestBody struct {

	// **参数解释**: 插件名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	PluginName string `json:"plugin_name"`

	// **参数解释**: 插件版本。 **约束限制**: 不涉及。 **取值范围**: 长度为2~64个字符。 **默认取值**: 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**: 插件安装端口。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	Port *string `json:"port,omitempty"`

	// **参数解释**: 安装插件的OBS地址。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	ObsUrl *string `json:"obs_url,omitempty"`
}

func (o InstallPluginRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstallPluginRequestBody struct{}"
	}

	return strings.Join([]string{"InstallPluginRequestBody", string(data)}, " ")
}
