package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type IssueDetailCustomFieldV2 struct {

	// **参数解释：** 自定义字段。 **取值范围：** 不涉及。
	CustomField *string `json:"custom_field,omitempty"`

	// **参数解释：** 自定义字段名称。 **取值范围：** 不涉及。
	FieldName *string `json:"field_name,omitempty"`

	// **参数解释：** 自定义属性对应的值，多个值以英文逗号区分开。 **取值范围：** 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释：** 自定义字段类型。 **取值范围：** textArea 多行文本 text 单行文本 select 下拉框 number 数字 time_date 日期 checkbox 多选框 radio 单选框。
	FieldType *string `json:"field_type,omitempty"`

	// **参数解释：** 自定义字段描述。 **取值范围：** 不涉及。
	Description *string `json:"description,omitempty"`
}

func (o IssueDetailCustomFieldV2) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "IssueDetailCustomFieldV2 struct{}"
	}

	return strings.Join([]string{"IssueDetailCustomFieldV2", string(data)}, " ")
}
