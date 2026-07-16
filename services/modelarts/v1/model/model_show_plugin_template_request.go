package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPluginTemplateRequest Request Object
type ShowPluginTemplateRequest struct {

	// **参数解释**：插件模板的名称。 **约束限制**：不涉及。 **取值范围**：可选值如下： - gpu-driver：GPU驱动插件模板信息 - npu-driver：NPU驱动插件模板信息 **默认取值**：不涉及。
	PlugintemplateName string `json:"plugintemplate_name"`
}

func (o ShowPluginTemplateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPluginTemplateRequest struct{}"
	}

	return strings.Join([]string{"ShowPluginTemplateRequest", string(data)}, " ")
}
