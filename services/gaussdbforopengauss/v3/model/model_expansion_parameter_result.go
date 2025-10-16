package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ExpansionParameterResult struct {

	// **参数解释**: 参数名称。 **取值范围**: 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**: 参数值。 **取值范围**: 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**: 修改参数是否需要重启。 **取值范围**: - true：需要重启。 - false：不需要重启。
	RestartRequired *bool `json:"restart_required,omitempty"`

	// **参数解释**: 参数值的取值范围。 **取值范围**: 不涉及。
	ValueRange *string `json:"value_range,omitempty"`

	// **参数解释**: 参数类型。 **取值范围**: - integer：整数。 - boolean：布尔类型。 - string：字符串。
	Type *string `json:"type,omitempty"`

	// **参数解释**: 参数解释。 **取值范围**: 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o ExpansionParameterResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExpansionParameterResult struct{}"
	}

	return strings.Join([]string{"ExpansionParameterResult", string(data)}, " ")
}
