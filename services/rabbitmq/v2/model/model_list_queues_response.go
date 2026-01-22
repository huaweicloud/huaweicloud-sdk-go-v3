package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListQueuesResponse Response Object
type ListQueuesResponse struct {

	// **参数解释**： 当前显示数量。 **取值范围**： 不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**： 查询结果总数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 查询详情。
	Items          *[]QueueDetails `json:"items,omitempty"`
	HttpStatusCode int             `json:"-"`
}

func (o ListQueuesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListQueuesResponse struct{}"
	}

	return strings.Join([]string{"ListQueuesResponse", string(data)}, " ")
}
