package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MarkerPageInfo **参数解释**：分页游标信息，包含上一页游标、下一页游标和当前页记录数。 **约束限制**：不涉及。
type MarkerPageInfo struct {

	// **参数解释**：下一页查询marker值，若为空表示当前已是最后一页。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	NextMarker *string `json:"next_marker,omitempty"`

	// **参数解释**：返回前一页查询地址，为空则表示是第一页。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	PreviousMarker *string `json:"previous_marker,omitempty"`

	// **参数解释**：本次查询记录总数。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	CurrentCount *int32 `json:"current_count,omitempty"`
}

func (o MarkerPageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MarkerPageInfo struct{}"
	}

	return strings.Join([]string{"MarkerPageInfo", string(data)}, " ")
}
