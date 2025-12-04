package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInstanceMessagesResponse Response Object
type ShowInstanceMessagesResponse struct {

	// **参数解释**： 消息列表。
	Messages *[]MessagesEntity `json:"messages,omitempty"`

	// **参数解释**： 消息总条数。 **取值范围**： 不涉及。
	Total *int64 `json:"total,omitempty"`

	// **参数解释**： 每页消息条数。 **取值范围**： 不涉及。
	Size           *int64 `json:"size,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowInstanceMessagesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceMessagesResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceMessagesResponse", string(data)}, " ")
}
