package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WorkItemLabelVo 标签对象
type WorkItemLabelVo struct {

	// **参数解释：**  标签Id。 **约束限制：**  最小长度：18，最大长度：19。 **取值范围：**  不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：**  对象类型。 **约束限制：**  固定为Label，表示当前对象类型为标签。 **取值范围：**  不涉及。
	Category *string `json:"category,omitempty"`

	// **参数解释：**  标签所属的工作项的类别。 **取值范围：**  - requirement - raw requirement - bug - task - feature
	LabelType *string `json:"label_type,omitempty"`

	// **参数解释：**  标签颜色的RGB值。 **取值范围：**  - #86CAFF - #6DDEBB - #A6DD82 - #FAC20A  - #FA9841 - #F66F6A - #F3689A - #A97AF8 - #71757F - #5E7CE0 - #207AB3 - #169E6C - #6CA83B - #B58200 - #B54E04 - #B02121 - #AD215B - #572DB3 - #4F4F4F - #3C51A6
	Color *string `json:"color,omitempty"`

	// **参数解释：**  标签标题。 **约束限制：**  最小长度：1，最大长度：30。
	Title *string `json:"title,omitempty"`

	// **参数解释：**  表示当前对象数据类型为标签。 **约束限制：**  固定为label。 **取值范围：**  label。
	Type *string `json:"type,omitempty"`

	// **参数解释：**  标签的生命周期。 **取值范围：**  - 正在工作 - 作废 - 删除
	State *string `json:"state,omitempty"`

	// **参数解释：**  最近修改人。 **约束限制：**  不涉及。
	ModifiedBy *string `json:"modified_by,omitempty"`
}

func (o WorkItemLabelVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WorkItemLabelVo struct{}"
	}

	return strings.Join([]string{"WorkItemLabelVo", string(data)}, " ")
}
