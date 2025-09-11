package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DimensionListResult struct {

	// **参数解释**: 统计数据名称。 **取值范围**: 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**: 活跃会话数。 **取值范围**: 不涉及。
	ActiveNum *string `json:"active_num,omitempty"`

	// **参数解释**: 总会话数。 **取值范围**: 不涉及。
	TotalNum *string `json:"total_num,omitempty"`
}

func (o DimensionListResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DimensionListResult struct{}"
	}

	return strings.Join([]string{"DimensionListResult", string(data)}, " ")
}
