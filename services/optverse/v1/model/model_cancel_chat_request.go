package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelChatRequest Request Object
type CancelChatRequest struct {

	// **参数解释**： 对话路由ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-64]个字符。 **默认取值**： 不涉及
	XChatRouteId string `json:"X-Chat-Route-Id"`

	// **参数解释**： 对话ID。 **约束限制**： 不涉及 **取值范围**： 仅支持字母、数字、中划线和下划线，长度为[1-128]个字符。 **默认取值**： 不涉及
	ChatId string `json:"chat_id"`
}

func (o CancelChatRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelChatRequest struct{}"
	}

	return strings.Join([]string{"CancelChatRequest", string(data)}, " ")
}
