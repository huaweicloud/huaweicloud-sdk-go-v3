package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateVariableGroupReqVariables struct {

	// **参数解释**： 参数序号。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Sequence *int32 `json:"sequence,omitempty"`

	// **参数解释**： 参数名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 参数类型。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 参数默认值。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**： 参数描述。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 是否私密参数。 **约束限制**： 不涉及。 **取值范围**： - true：是私密参数。 - false：不是私密参数。 **默认取值**： 不涉及。
	IsSecret *bool `json:"is_secret,omitempty"`
}

func (o CreateVariableGroupReqVariables) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVariableGroupReqVariables struct{}"
	}

	return strings.Join([]string{"CreateVariableGroupReqVariables", string(data)}, " ")
}
