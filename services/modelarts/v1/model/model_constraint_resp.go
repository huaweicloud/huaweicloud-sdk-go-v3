package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConstraintResp 条件。
type ConstraintResp struct {

	// **参数解释**：条件属性，参数的某个字段值。 **取值范围**：不涉及。
	Attribute *string `json:"attribute,omitempty"`

	// **参数解释**：操作。 **取值范围**：不涉及。
	Operator *string `json:"operator,omitempty"`

	// **参数解释**：取值。
	Value *interface{} `json:"value,omitempty"`
}

func (o ConstraintResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConstraintResp struct{}"
	}

	return strings.Join([]string{"ConstraintResp", string(data)}, " ")
}
