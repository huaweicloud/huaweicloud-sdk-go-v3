package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowMessagesResponse Response Object
type ShowMessagesResponse struct {

	// **参数解释**： 消息列表。
	Messages *[]ShowMessagesRespMessages `json:"messages,omitempty"`

	// **参数解释**： 消息总数。 **取值范围**： 不涉及。
	MessagesCount *int32 `json:"messages_count,omitempty"`

	// **参数解释**： 总页数。 **取值范围**： 不涉及。
	OffsetsCount *int32 `json:"offsets_count,omitempty"`

	// **参数解释**： 当前页数。 **取值范围**： 不涉及。
	Offset         *int32 `json:"offset,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMessagesResponse struct{}"
	}

	return strings.Join([]string{"ShowMessagesResponse", string(data)}, " ")
}
