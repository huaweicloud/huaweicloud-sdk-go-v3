package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAccessPointResponseBodyPageInfo **参数解释：** 分页信息 **取值范围：** 不涉及
type ListAccessPointResponseBodyPageInfo struct {

	// **参数解释：** 下一页的marker **取值范围：** 不涉及
	NextMarker string `json:"next_marker"`

	// **参数解释：** 本页数量 **取值范围：** 不涉及
	CurrentCount int32 `json:"current_count"`
}

func (o ListAccessPointResponseBodyPageInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAccessPointResponseBodyPageInfo struct{}"
	}

	return strings.Join([]string{"ListAccessPointResponseBodyPageInfo", string(data)}, " ")
}
