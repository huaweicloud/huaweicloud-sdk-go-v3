package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Conditions struct {

	// **参数解释**：操作属性。 **取值范围**：不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**：操作符号。 **取值范围**：不涉及。
	Operator *string `json:"operator,omitempty"`

	// **参数解释**：操作符的值。
	Value *[]string `json:"value,omitempty"`
}

func (o Conditions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Conditions struct{}"
	}

	return strings.Join([]string{"Conditions", string(data)}, " ")
}
