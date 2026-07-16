package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListInferServiceEventsResponse Response Object
type ListInferServiceEventsResponse struct {

	// **参数解释：** 当前页。 **取值范围：** 不涉及。
	Current *int32 `json:"current,omitempty"`

	// **参数解释：** 总页数。 **取值范围：** 不涉及。
	Pages *int32 `json:"pages,omitempty"`

	// **参数解释：** 每一页的数量。 **取值范围：** 不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释：** 总记录数。 **取值范围：** 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释：** 服务事件列表。 **取值范围：** 不涉及。
	Data           *[]ServiceEventResponse `json:"data,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ListInferServiceEventsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInferServiceEventsResponse struct{}"
	}

	return strings.Join([]string{"ListInferServiceEventsResponse", string(data)}, " ")
}
