package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CategoryLayerDto Category层级数据类型
type CategoryLayerDto struct {
	Category *BaseCategory `json:"category,omitempty"`

	// **参数解释**： 父类字段。 **取值范围**： 不涉及。
	LinkParentField *string `json:"link_parent_field,omitempty"`

	// **参数解释**： 工作项层级ID。 **取值范围**： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释**： 子工作项层级数据类型。 **取值范围**： 不涉及。
	Children *[]CategoryLayerDto `json:"children,omitempty"`

	// **参数解释**： 层级对象类型编码。 **取值范围**： 不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释**： 对象类型编码。 **取值范围**： 不涉及。
	CategoryCode *string `json:"category_code,omitempty"`

	// **参数解释**： 对象类型ID。 **取值范围**： 不涉及。
	CategoryId *string `json:"category_id,omitempty"`

	// **参数解释**： 层级类型。 **取值范围**： 不涉及。
	LayerType *string `json:"layer_type,omitempty"`

	// **参数解释**： 父ID。 **取值范围**： 不涉及。
	ParentId *string `json:"parent_id,omitempty"`

	// **参数解释**： 根工作项ID。 **取值范围**： 不涉及。
	RootId *string `json:"root_id,omitempty"`

	// **参数解释**： 画布X轴坐标。 **取值范围**： 不涉及。
	PositionX *int32 `json:"position_x,omitempty"`

	// **参数解释**： 画布Y轴坐标。 **取值范围**： 不涉及。
	PositionY *int32 `json:"position_y,omitempty"`
}

func (o CategoryLayerDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CategoryLayerDto struct{}"
	}

	return strings.Join([]string{"CategoryLayerDto", string(data)}, " ")
}
