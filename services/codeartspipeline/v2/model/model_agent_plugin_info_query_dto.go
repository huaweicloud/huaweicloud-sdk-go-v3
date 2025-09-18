package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AgentPluginInfoQueryDto struct {

	// **参数解释**： 可选的查询条件-插件名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 可选的查询条件-匹配名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RegexName *string `json:"regex_name,omitempty"`

	// **参数解释**： 维护者。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Maintainer *string `json:"maintainer,omitempty"`

	// **参数解释**： 业务类型。 **约束限制**： 不涉及。 **取值范围**： - Build。 - Gate。 - Deploy。 - Test。 - Normal。 **默认取值**： 不涉及。
	BusinessType *[]string `json:"business_type,omitempty"`

	// **参数解释**： 插件属性。 **约束限制**： 不涉及。 **取值范围**： - official：自定义插件。 - custom：基础插件。 **默认取值**： 不涉及。
	PluginAttribution *string `json:"plugin_attribution,omitempty"`
}

func (o AgentPluginInfoQueryDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AgentPluginInfoQueryDto struct{}"
	}

	return strings.Join([]string{"AgentPluginInfoQueryDto", string(data)}, " ")
}
