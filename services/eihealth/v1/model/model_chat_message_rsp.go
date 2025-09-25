package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChatMessageRsp struct {

	// **参数解释**： 问答ID。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	Id *string `json:"id,omitempty"`

	Role *MessageRole `json:"role,omitempty"`

	Type *QaType `json:"type,omitempty"`

	// **参数解释**： 问答创建时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	CreateTime *string `json:"create_time,omitempty"`

	// **参数解释**： 问答更新时间。 **约束限制**： 不涉及 **取值范围**： 不涉及 **默认取值**： 不涉及
	UpdateTime *string `json:"update_time,omitempty"`

	Message *ChatMessage `json:"message,omitempty"`

	Addition *ChatAddition `json:"addition,omitempty"`
}

func (o ChatMessageRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChatMessageRsp struct{}"
	}

	return strings.Join([]string{"ChatMessageRsp", string(data)}, " ")
}
