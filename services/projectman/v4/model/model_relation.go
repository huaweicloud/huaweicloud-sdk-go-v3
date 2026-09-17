package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Relation 关联关系
type Relation struct {

	// **参数解释**： 关系code。 **取值范围**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 关联的工作项code列表。 **取值范围**： 不涉及。
	Categories *[]string `json:"categories,omitempty"`

	// **参数解释**： 工作流场景使用，前置校验中的关联关系校验字段。 **取值范围**： 不涉及。
	LinkFieldCode *string `json:"link_field_code,omitempty"`

	// **参数解释**： 关系名称，在工作项详情关联项下左侧显示。 **取值范围**： 不涉及。
	RelationName *string `json:"relation_name,omitempty"`

	// **参数解释**： 关系描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 展示范围。 **取值范围**： 不涉及。
	DisplayScope *string `json:"display_scope,omitempty"`

	// **参数解释**： 动作行为。 **取值范围**： 不涉及。
	Actions *[]RelateAction `json:"actions,omitempty"`

	// **参数解释**： 动作行为。 **取值范围**： - ONE_TO_ONE 一对一 - ONE_TO_MANY 一对多 - MANY_TO_ONE 多对一 - MANY_TO_MANY 多对多
	RelateType *string `json:"relate_type,omitempty"`
}

func (o Relation) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Relation struct{}"
	}

	return strings.Join([]string{"Relation", string(data)}, " ")
}
