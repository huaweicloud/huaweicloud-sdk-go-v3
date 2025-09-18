package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateRuleReq struct {

	// **参数解释**： 创建规则的类型。 **约束限制**： 不涉及。 **取值范围**： - Build。 - Gate。 - Deploy。 - Test。 - Normal。 **默认取值**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 规则名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 布局内容。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	LayoutContent string `json:"layout_content"`

	// **参数解释**： 扩展插件唯一ID。可以通过[查询插件版本详情](ShowPluginVersion.xml)接口，获取响应参数中unique_id。 **约束限制**： 不涉及。 **取值范围**： 32位字符，由数字和字母组成。 **默认取值**： 不涉及。
	PluginId *string `json:"plugin_id,omitempty"`

	// **参数解释**： 扩展插件名称。 **约束限制**： 仅支持输入大小写英文字母、数字、'-'、'_'。 **取值范围**： 1到50位字符。 **默认取值**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 插件版本号。 **约束限制**： 必须是类似 x.xx.xx 的格式，其中：x 是 1 到 2 位的数字（范围 0 到 99）。xx 是点后跟随的数字部分，且每部分可以是 1 位或 2 位数字。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PluginVersion *string `json:"plugin_version,omitempty"`

	// **参数解释**： 规则属性分组列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Content []RuleContent `json:"content"`
}

func (o CreateRuleReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateRuleReq struct{}"
	}

	return strings.Join([]string{"CreateRuleReq", string(data)}, " ")
}
