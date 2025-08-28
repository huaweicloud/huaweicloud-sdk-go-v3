package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopIpResponse Response Object
type ListTopIpResponse struct {

	// **参数解释：** CountItem的总数量 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** CountItem详细信息列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items          *[]CountItem `json:"items,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListTopIpResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopIpResponse struct{}"
	}

	return strings.Join([]string{"ListTopIpResponse", string(data)}, " ")
}
