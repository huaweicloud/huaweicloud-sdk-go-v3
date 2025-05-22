package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ActionItemVo **参数解释**： 任务详情子项。 **取值范围**： 不涉及。
type ActionItemVo struct {

	// **参数解释**： 任务详情子项，一级菜单。 **取值范围**： 不涉及。
	ItemName *string `json:"item_name,omitempty"`

	// **参数解释**： 任务详情子项，一级菜单详情。 **取值范围**： 不涉及。
	SubItems *[]ActionSubItemVo `json:"sub_items,omitempty"`
}

func (o ActionItemVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionItemVo struct{}"
	}

	return strings.Join([]string{"ActionItemVo", string(data)}, " ")
}
