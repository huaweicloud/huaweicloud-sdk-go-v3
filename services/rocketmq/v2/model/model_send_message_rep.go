package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SendMessageRep struct {

	// **参数解释**： 主题名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 消息内容。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Body *string `json:"body,omitempty"`

	// **参数解释**： 特性列表。
	PropertyList *[]SendMessageProperties `json:"property_list,omitempty"`
}

func (o SendMessageRep) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SendMessageRep struct{}"
	}

	return strings.Join([]string{"SendMessageRep", string(data)}, " ")
}
