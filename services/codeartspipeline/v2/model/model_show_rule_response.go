package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRuleResponse Response Object
type ShowRuleResponse struct {

	// **参数解释**： 规则ID。 **取值范围**： 32位字符，由数字和字母组成。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 规则类型。 **取值范围**： - Build：构建。 - Gate：代码检查。 - Test：测试。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 规则名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 规则是否有效。 **取值范围**： true:有效，false:无效
	IsValid *bool `json:"is_valid,omitempty"`

	// **参数解释**： 规则版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 扩展插件唯一ID。 **取值范围**： 不涉及。
	PluginId *string `json:"plugin_id,omitempty"`

	// **参数解释**： 扩展插件名称。 **取值范围**： 不涉及。
	PluginName *string `json:"plugin_name,omitempty"`

	// **参数解释**： 扩展插件版本号。 **取值范围**： 不涉及。
	PluginVersion *string `json:"plugin_version,omitempty"`

	// **参数解释**： 规则创建人。 **取值范围**： 不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释**： 规则创建时间。 **取值范围**： 不涉及。
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**： 规则最后更新人。 **取值范围**： 不涉及。
	Updater *string `json:"updater,omitempty"`

	// **参数解释**： 规则最后更新时间。 **取值范围**： 不涉及。
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 规则详细属性。 **取值范围**： 不涉及。
	Content        *[]RuleContent `json:"content,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowRuleResponse", string(data)}, " ")
}
