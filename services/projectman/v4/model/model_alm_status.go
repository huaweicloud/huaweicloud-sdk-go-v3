package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlmStatus struct {

	// **参数解释：**  状态Id。 **取值范围：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  工作项的状态属性。 **取值范围：**  - START - IN_PROGRESS - END
	Belonging *string `json:"belonging,omitempty"`

	// **参数解释：**  状态所属的项目空间id。 **取值范围：**  不涉及。
	SpaceId *string `json:"space_id,omitempty"`

	// **参数解释：**  状态名称。 **取值范围：**  不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：**  状态code值。 **取值范围：**  不涉及。
	Code *string `json:"code,omitempty"`

	// **参数解释：**  状态定义级别，1,2,3为系统级，4为租户自定义，5为项目自定义。 **取值范围：**  不涉及。
	DefinitionType *string `json:"definition_type,omitempty"`

	// **参数解释：**  状态归属定义级别，1,2,3为系统级，4为租户自定义，5为项目自定义。区别于definition_type。如果为系统级和租户自定义级，在项目中会复制一份元数据，归属于项目空间。 **取值范围：**  不涉及。
	BelongDefinitionType *int32 `json:"belong_definition_type,omitempty"`

	// **参数解释：**  状态名称，和name值相同。 **取值范围：**  不涉及。
	DisplayValue *string `json:"display_value,omitempty"`

	// **参数解释：**  位置顺序。 **取值范围：**  不涉及。
	Position *int32 `json:"position,omitempty"`

	// **参数解释：**  是否显示。 **取值范围：**  不涉及。
	Displayable *int32 `json:"displayable,omitempty"`

	// **参数解释：**  是否可编辑。 **取值范围：**  不涉及。
	Editable *int32 `json:"editable,omitempty"`

	// **参数解释：**  是否可删除。 **取值范围：**  不涉及。
	Deletable *int32 `json:"deletable,omitempty"`

	// **参数解释：**  是否可变，即是否为固定值。 **取值范围：**  不涉及。
	Mutable *int32 `json:"mutable,omitempty"`

	// **参数解释：**  标题的拼音首字母。 **取值范围：**  不涉及。
	TitlePy *string `json:"title_py,omitempty"`

	// **参数解释：**  创建人用户Id。 **取值范围：**  不涉及。
	CreatedBy *string `json:"created_by,omitempty"`

	// **参数解释：**  创建时间。Unix时间戳，精度为毫秒。 **取值范围：**  不涉及。
	CreatedDate *int64 `json:"created_date,omitempty"`

	// **参数解释：**  最近修改时间。Unix时间戳，精度为毫秒。 **取值范围：**  不涉及。
	ModifiedDate *int64 `json:"modified_date,omitempty"`

	// **参数解释：**  最近修改人用户Id。 **取值范围：**    不涉及。
	ModifiedBy *string `json:"modified_by,omitempty"`
}

func (o AlmStatus) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlmStatus struct{}"
	}

	return strings.Join([]string{"AlmStatus", string(data)}, " ")
}
