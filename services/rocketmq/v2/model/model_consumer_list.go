package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerList struct {

	// **参数解释**： Topic列表（当查询Topic消费“列表”时才显示此参数）。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Topics *[]string `json:"topics,omitempty"`

	// **参数解释**： Topic总数（当查询Topic消费“列表”时才显示此参数）。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Total *int32 `json:"total,omitempty"`
}

func (o ConsumerList) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerList struct{}"
	}

	return strings.Join([]string{"ConsumerList", string(data)}, " ")
}
