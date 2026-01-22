package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeadletterResendRespResendResults struct {

	// **参数解释**： 消息ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	MsgId *string `json:"msg_id,omitempty"`

	// **参数解释**： 错误码。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	ErrorCode *string `json:"error_code,omitempty"`

	// **参数解释**： 错误信息。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	ErrorMessage *string `json:"error_message,omitempty"`
}

func (o DeadletterResendRespResendResults) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadletterResendRespResendResults struct{}"
	}

	return strings.Join([]string{"DeadletterResendRespResendResults", string(data)}, " ")
}
