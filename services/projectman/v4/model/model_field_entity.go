package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// FieldEntity 字段对象
type FieldEntity struct {

	// **参数解释**： 字段名称。 **取值范围**： 不涉及
	DisplayName *string `json:"display_name,omitempty"`

	// **参数解释**： 字段code。 **取值范围**： 不涉及
	Code *string `json:"code,omitempty"`

	// **参数解释**： 字段id。 **取值范围**： 不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**： 字段描述。 **取值范围**： 不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**： 字段创建人名称。 **取值范围**： 不涉及
	CreatedBy *string `json:"created_by,omitempty"`

	// **参数解释**： 字段创建时间，时间戳格式，示例:1715305846000。 **取值范围**： 不涉及
	CreatedDate *int64 `json:"created_date,omitempty"`

	// **参数解释**： 字段最后更新人名称。 **取值范围**： 不涉及
	ModifiedBy *string `json:"modified_by,omitempty"`

	// **参数解释**： 字段级别。 **取值范围**： 1/2/3：系统预设字段。 4：租户自定义字段 5：项目自定义字段
	DefinitionType *string `json:"definition_type,omitempty"`

	// **参数解释**： 字段类型名称。 **取值范围**： 不涉及
	FieldTypeName *string `json:"field_type_name,omitempty"`

	// **参数解释**： 字段在工作项中是否必填，和工作流配置不一致时以工作流为准。 **取值范围**： true（必填） false（非必填）
	Required *bool `json:"required,omitempty"`

	// **参数解释**： 字段在工作项中是否受控，修改已基线的工作项受控字段需要走变更评审流程，和工作流配置不一致时以工作流为准。 **取值范围**： true（受控） false（非受控）
	Controlled *bool `json:"controlled,omitempty"`

	// **参数解释**： 字段在工作项中是否可修改，和工作流配置不一致时以工作流为准。 **取值范围**： true（不可修改） false（可修改）
	Immutable *bool `json:"immutable,omitempty"`

	// **参数解释**： 字段排序的序号。 **取值范围**： 不涉及
	No *int32 `json:"no,omitempty"`

	// **参数解释**： 字段选项。
	AllOptions *[]OptionEntity `json:"all_options,omitempty"`
}

func (o FieldEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "FieldEntity struct{}"
	}

	return strings.Join([]string{"FieldEntity", string(data)}, " ")
}
