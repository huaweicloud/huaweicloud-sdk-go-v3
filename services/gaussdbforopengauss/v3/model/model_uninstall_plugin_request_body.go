package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UninstallPluginRequestBody struct {

	// **参数解释**: 插件名称。 **约束限制**: 不涉及。 **取值范围**: 不涉及。 **默认取值**: 不涉及。
	PluginName string `json:"plugin_name"`
}

func (o UninstallPluginRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UninstallPluginRequestBody struct{}"
	}

	return strings.Join([]string{"UninstallPluginRequestBody", string(data)}, " ")
}
