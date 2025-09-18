package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowStrategyResponse Response Object
type ShowStrategyResponse struct {

	// **参数解释**： 规则模板实例ID。 **取值范围**： 32位字符，由数字和字母组成。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 规则模板实例名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 规则模板类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 规则模板策略。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 策略创建人。 **取值范围**： 不涉及。
	Creator *string `json:"creator,omitempty"`

	// **参数解释**： 策略创建时间。 **取值范围**： 不涉及。
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**： 策略最近更新人。 **取值范围**： 不涉及。
	Updater *string `json:"updater,omitempty"`

	// **参数解释**： 策略最近更新时间。 **取值范围**： 不涉及。
	UpdateTime *string `json:"update_time,omitempty"`

	// **参数解释**： 策略是否生效。 **取值范围**： - true：策略生效。 - false：策略不生效。
	IsValid *bool `json:"is_valid,omitempty"`

	// **参数解释**： 规则实例集合。 **取值范围**： 不涉及。
	RuleInstances *[]RuleInstance `json:"rule_instances,omitempty"`

	// **参数解释**： 规则实例生效级别。 **取值范围**： 不涉及。
	Level *string `json:"level,omitempty"`

	// **参数解释**： 规则实例是否系统级。 **取值范围**： - true：规则实例是系统级。 - false：规则实例不是系统级。
	IsPublic *bool `json:"is_public,omitempty"`

	// **参数解释**： 规则实例是1.0的数据。 **取值范围**： - true：规则实例是1.0的数据。 - false：规则实例是1.0的数据。
	IsLegacy       *bool `json:"is_legacy,omitempty"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowStrategyResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowStrategyResponse struct{}"
	}

	return strings.Join([]string{"ShowStrategyResponse", string(data)}, " ")
}
