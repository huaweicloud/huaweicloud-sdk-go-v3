package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type DeadletterResendReq struct {

	// **参数解释**： Topic名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 消息列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	MsgIdList *[]string `json:"msg_id_list,omitempty"`
}

func (o DeadletterResendReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeadletterResendReq struct{}"
	}

	return strings.Join([]string{"DeadletterResendReq", string(data)}, " ")
}
