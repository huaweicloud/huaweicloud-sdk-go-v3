package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumersInTagEntity struct {

	// **参数解释**： 消费者列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Consumers *[]string `json:"consumers,omitempty"`

	// **参数解释**： 标签名。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TagName *string `json:"tag_name,omitempty"`
}

func (o ConsumersInTagEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumersInTagEntity struct{}"
	}

	return strings.Join([]string{"ConsumersInTagEntity", string(data)}, " ")
}
