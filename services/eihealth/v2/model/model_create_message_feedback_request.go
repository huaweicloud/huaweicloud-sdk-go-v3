package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateMessageFeedbackRequest Request Object
type CreateMessageFeedbackRequest struct {

	// **参数解释**： 对话ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	ChatId string `json:"chat_id"`

	// **参数解释**： 问答ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	MessageId string `json:"message_id"`

	Body *CreateMessageFeedbackReq `json:"body,omitempty"`
}

func (o CreateMessageFeedbackRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateMessageFeedbackRequest struct{}"
	}

	return strings.Join([]string{"CreateMessageFeedbackRequest", string(data)}, " ")
}
