package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ActionDistributions struct {

	// **参数解释：** 处置动作的总数量 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Total *int64 `json:"total,omitempty"`

	// **参数解释：** 各处置动作的统计详情 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items *[]TypedStatBucket `json:"items,omitempty"`
}

func (o ActionDistributions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ActionDistributions struct{}"
	}

	return strings.Join([]string{"ActionDistributions", string(data)}, " ")
}
