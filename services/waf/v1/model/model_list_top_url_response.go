package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTopUrlResponse Response Object
type ListTopUrlResponse struct {

	// **参数解释：** CountItem的总数量 **约束限制：** 不涉及 **取值范围：** ≥0 **默认取值：** 0
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** CountItem详细信息列表 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** 不涉及
	Items          *[]UrlCountItem `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListTopUrlResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTopUrlResponse struct{}"
	}

	return strings.Join([]string{"ListTopUrlResponse", string(data)}, " ")
}
