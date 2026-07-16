package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPluginTemplatesRequest Request Object
type ListPluginTemplatesRequest struct {

	// **参数解释**：指定的插件名称，填写则查询指定名称的插件。 **约束限制**：不涉及 **取值范围**：不涉及。 **默认取值**：不涉及。
	TemplateName *string `json:"templateName,omitempty"`

	// **参数解释**：指定的资源池名称，填写则查询符合资源池安装条件的插件列表。 **约束限制**：不涉及 **取值范围**：不涉及。 **默认取值**：不涉及。
	PoolName *string `json:"poolName,omitempty"`
}

func (o ListPluginTemplatesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPluginTemplatesRequest struct{}"
	}

	return strings.Join([]string{"ListPluginTemplatesRequest", string(data)}, " ")
}
