package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SendMessageReq struct {

	// **参数解释**： 主题名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topic string `json:"topic"`

	// **参数解释**： 消息内容。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Body string `json:"body"`

	// **参数解释**： 特性列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值** 不涉及。
	PropertyList *[]SendMessageProperties `json:"property_list,omitempty"`
}

func (o SendMessageReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendMessageReq struct{}"
	}

	return strings.Join([]string{"SendMessageReq", string(data)}, " ")
}
