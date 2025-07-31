package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PageInfo **参数解释：** 分页信息。 **取值范围：** 不涉及。
type PageInfo struct {

	// **参数解释：** 下一页的页面标识。 **取值范围：** 不涉及。
	NextMarker *string `json:"next_marker,omitempty"`

	// **参数解释：** 上一页的页面标识。 **取值范围：** 不涉及。
	PreviousMarker *string `json:"previous_marker,omitempty"`

	// **参数解释：** 页面数量。 **取值范围：** 不涉及。
	CurrentCount *int32 `json:"current_count,omitempty"`
}

func (o PageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfo struct{}"
	}

	return strings.Join([]string{"PageInfo", string(data)}, " ")
}
