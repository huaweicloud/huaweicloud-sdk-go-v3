package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PlanListResponsePage **参数解释**： 计划列表分页信息。
type PlanListResponsePage struct {

	// **参数解释：** 页码 **取值范围：** 不涉及
	Page *int32 `json:"page,omitempty"`

	// **参数解释：** 分页数量 **取值范围：** 不涉及
	Size *int32 `json:"size,omitempty"`

	// **参数解释：** 当前页数量 **取值范围：** 不涉及
	Count *int32 `json:"count,omitempty"`
}

func (o PlanListResponsePage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PlanListResponsePage struct{}"
	}

	return strings.Join([]string{"PlanListResponsePage", string(data)}, " ")
}
