package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ResendReq struct {

	// **参数解释**： Group ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	Group string `json:"group"`

	// **参数解释**： 消息所属Topic。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 客户端ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	ClientId string `json:"client_id"`

	// **参数解释**： 消息列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	MsgIdList *[]string `json:"msg_id_list,omitempty"`
}

func (o ResendReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ResendReq struct{}"
	}

	return strings.Join([]string{"ResendReq", string(data)}, " ")
}
