package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListMessagesForRocketMqResponse Response Object
type ListMessagesForRocketMqResponse struct {

	// **参数解释**： 消息列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Messages *[]Message `json:"messages,omitempty"`

	// **参数解释**： 消息总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Total          float32 `json:"total,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ListMessagesForRocketMqResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListMessagesForRocketMqResponse struct{}"
	}

	return strings.Join([]string{"ListMessagesForRocketMqResponse", string(data)}, " ")
}
