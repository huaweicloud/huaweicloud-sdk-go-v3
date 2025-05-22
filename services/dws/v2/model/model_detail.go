package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Detail 规格详细信息。
type Detail struct {

	// **参数解释**： 属性类型。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 属性值。 **取值范围**： 不涉及。
	Value string `json:"value"`

	// **参数解释**： 属性单位。 **取值范围**： 不涉及。
	Unit *string `json:"unit,omitempty"`
}

func (o Detail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Detail struct{}"
	}

	return strings.Join([]string{"Detail", string(data)}, " ")
}
