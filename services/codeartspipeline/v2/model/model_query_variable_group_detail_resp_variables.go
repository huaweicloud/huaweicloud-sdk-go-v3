package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type QueryVariableGroupDetailRespVariables struct {

	// **参数解释**： 参数名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 参数序号。 **取值范围**： 不涉及。
	Sequence *int32 `json:"sequence,omitempty"`

	// **参数解释**： 参数类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 参数默认值。 **取值范围**： 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**： 是否私密。 **取值范围**： - true：是私密参数。 - false：不是私密参数。
	IsSecret *bool `json:"is_secret,omitempty"`

	// **参数解释**： 描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o QueryVariableGroupDetailRespVariables) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "QueryVariableGroupDetailRespVariables struct{}"
	}

	return strings.Join([]string{"QueryVariableGroupDetailRespVariables", string(data)}, " ")
}
