package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleInstance struct {

	// **参数解释**： 规则实例ID。 **取值范围**： 32位字符，由数字和字母组成。
	Id string `json:"id"`

	// **参数解释**： 规则类型。 **取值范围**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 规则名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 规则版本。 **取值范围**： 不涉及。
	Version string `json:"version"`

	// **参数解释**： 插件ID。 **取值范围**： 不涉及。
	PluginId *string `json:"plugin_id,omitempty"`

	// **参数解释**： 插件名称。 **取值范围**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 插件版本号。 **取值范围**： 不涉及。
	PluginVersion *string `json:"plugin_version,omitempty"`

	// **参数解释**： 规则是否生效。 **取值范围**： - true：规则生效。 - false：规则不生效。
	IsValid bool `json:"is_valid"`

	// **参数解释**： 规则是否可编辑。 **取值范围**： - true：规则可编辑。 - false：规则不可编辑。
	Editable *bool `json:"editable,omitempty"`

	// **参数解释**： 规则实例集合。 **取值范围**： 不涉及。
	Content []RuleInstanceContent `json:"content"`

	Parent *RuleSet `json:"parent,omitempty"`
}

func (o RuleInstance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleInstance struct{}"
	}

	return strings.Join([]string{"RuleInstance", string(data)}, " ")
}
