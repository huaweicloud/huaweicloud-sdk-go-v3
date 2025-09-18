package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RuleSet struct {

	// **参数解释**： 规则模板实例ID。 **取值范围**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 规则模板实例名称。 **取值范围**： 不涉及。
	Name string `json:"name"`

	// **参数解释**： 规则实例类型。 **取值范围**： 不涉及。
	Type string `json:"type"`

	// **参数解释**： 规则实例版本。 **取值范围**： 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释**： 规则实例最近操作人。 **取值范围**： 不涉及。
	Operator string `json:"operator"`

	// **参数解释**： 规则实例最近操作时间。 **取值范围**： 不涉及。
	OperateTime int64 `json:"operate_time"`

	// **参数解释**： 规则实例是否生效。 **取值范围**： - true：规则生效。 - false：规则不生效。
	IsValid bool `json:"is_valid"`

	// **参数解释**： 规则实例生效级别。 **取值范围**： 不涉及。
	Level *string `json:"level,omitempty"`

	// **参数解释**： 规则实例是否系统级。 **取值范围**： - true：规则实例是系统级。 - false：规则实例不是系统级。
	IsPublic *bool `json:"is_public,omitempty"`

	// **参数解释**： 规则实例是1.0的数据。 **取值范围**： - true：规则实例是1.0的数据。 - false：规则实例是1.0的数据。
	IsLegacy *bool `json:"is_legacy,omitempty"`
}

func (o RuleSet) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RuleSet struct{}"
	}

	return strings.Join([]string{"RuleSet", string(data)}, " ")
}
