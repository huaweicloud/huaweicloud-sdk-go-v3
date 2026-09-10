package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListChatResponse Response Object
type ListChatResponse struct {

	// **参数解释**： 对话列表。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Chats *[]ChatRsp `json:"chats,omitempty"`

	// **参数解释**： 对话个数。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Count          *int32 `json:"count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ListChatResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListChatResponse struct{}"
	}

	return strings.Join([]string{"ListChatResponse", string(data)}, " ")
}
