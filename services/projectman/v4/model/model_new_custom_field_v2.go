package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NewCustomFieldV2 自定义字段
type NewCustomFieldV2 struct {

	// **参数解释：** 自定义字段。 **取值范围：** 不涉及。
	CustomField *string `json:"custom_field,omitempty"`

	// **参数解释：** 自定义字段名称。 **取值范围：** 不涉及。
	FieldName *string `json:"field_name,omitempty"`

	// **参数解释：** 自定义属性对应的值，多个值以英文逗号区分开。 **取值范围：** 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o NewCustomFieldV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NewCustomFieldV2 struct{}"
	}

	return strings.Join([]string{"NewCustomFieldV2", string(data)}, " ")
}
