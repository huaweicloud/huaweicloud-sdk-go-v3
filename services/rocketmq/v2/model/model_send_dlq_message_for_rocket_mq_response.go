package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// SendDlqMessageForRocketMqResponse Response Object
type SendDlqMessageForRocketMqResponse struct {

	// **参数解释**： 重发死信消息结果。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	ResendResults  *[]DeadletterResendRespResendResults `json:"resend_results,omitempty"`
	HttpStatusCode int                                  `json:"-"`
}

func (o SendDlqMessageForRocketMqResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendDlqMessageForRocketMqResponse struct{}"
	}

	return strings.Join([]string{"SendDlqMessageForRocketMqResponse", string(data)}, " ")
}
