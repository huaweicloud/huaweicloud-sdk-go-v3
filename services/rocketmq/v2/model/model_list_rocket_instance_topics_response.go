package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRocketInstanceTopicsResponse Response Object
type ListRocketInstanceTopicsResponse struct {

	// **参数解释**： Topic总数。 **取值范围**： 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**： 最大可创建Topic数量。 **取值范围**： 不涉及。
	Max *int32 `json:"max,omitempty"`

	// **参数解释**： 剩余可创建Topic数量。 **取值范围**： 不涉及。
	Remaining *int32 `json:"remaining,omitempty"`

	// **参数解释**： 下个分页的offset。 **取值范围**： 不涉及。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// **参数解释**： 上个分页的offset。 **取值范围**： 不涉及。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// **参数解释**： 剩余可创建Topic数量。
	Topics         *[]Topic `json:"topics,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListRocketInstanceTopicsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRocketInstanceTopicsResponse struct{}"
	}

	return strings.Join([]string{"ListRocketInstanceTopicsResponse", string(data)}, " ")
}
