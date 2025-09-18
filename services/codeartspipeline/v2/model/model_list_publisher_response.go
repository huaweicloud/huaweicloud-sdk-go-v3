package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPublisherResponse Response Object
type ListPublisherResponse struct {

	// **参数解释**： 查询偏移量，与查询参数的偏移量相同。 **取值范围**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 每次查询的条目数量。 **取值范围**： 大于等于0。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 本次查询到的数据条数。 **取值范围**： 不涉及。
	Total *int64 `json:"total,omitempty"`

	// **参数解释**： 本次查询到的数据列表。 **取值范围**： 不涉及。
	Data           *[]PublisherVo `json:"data,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListPublisherResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPublisherResponse struct{}"
	}

	return strings.Join([]string{"ListPublisherResponse", string(data)}, " ")
}
