package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMessageTraceResponse Response Object
type ListMessageTraceResponse struct {

	// **参数解释**： 总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Total float32 `json:"total,omitempty"`

	// **参数解释**： 下个分页的offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NextOffset *int32 `json:"next_offset,omitempty"`

	// **参数解释**： 上个分页的offset。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	PreviousOffset *int32 `json:"previous_offset,omitempty"`

	// **参数解释**： 消息轨迹列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Trace          *[]ListMessageTraceRespTrace `json:"trace,omitempty"`
	HttpStatusCode int                          `json:"-"`
}

func (o ListMessageTraceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessageTraceResponse struct{}"
	}

	return strings.Join([]string{"ListMessageTraceResponse", string(data)}, " ")
}
