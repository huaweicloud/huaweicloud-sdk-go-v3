package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CustomFieldV2 自定义属性。
type CustomFieldV2 struct {

	// **参数解释：** 自定义字段。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 自定义字段对应的值。 **取值范围：** 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释：** 自定义字段修改后的名称。 **取值范围：** 不涉及。
	NewName *string `json:"new_name,omitempty"`
}

func (o CustomFieldV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CustomFieldV2 struct{}"
	}

	return strings.Join([]string{"CustomFieldV2", string(data)}, " ")
}
