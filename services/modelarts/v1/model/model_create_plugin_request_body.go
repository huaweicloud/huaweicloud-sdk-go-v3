package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePluginRequestBody 创建插件实例的请求体。
type CreatePluginRequestBody struct {

	// **参数解释**：API资源类型，固定值“Plugin”，该值不可修改。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	ApiVersion string `json:"apiVersion"`

	// **参数解释**：API版本，固定值“v2”，该值不可修改。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Kind string `json:"kind"`

	Spec *PluginSpec `json:"spec"`
}

func (o CreatePluginRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePluginRequestBody struct{}"
	}

	return strings.Join([]string{"CreatePluginRequestBody", string(data)}, " ")
}
