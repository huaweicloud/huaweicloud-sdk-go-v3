package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ConsumerDetailResp struct {

	// **参数解释**： 消费堆积总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Lag *int64 `json:"lag,omitempty"`

	// **参数解释**： 消息总数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	MaxOffset *int64 `json:"max_offset,omitempty"`

	// **参数解释**： 已消费消息数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	ConsumerOffset *int64 `json:"consumer_offset,omitempty"`

	// **参数解释**： Topic关联代理（当查询Topic消费“详情”才显示此参数）。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Brokers *[]Brokers `json:"brokers,omitempty"`
}

func (o ConsumerDetailResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConsumerDetailResp struct{}"
	}

	return strings.Join([]string{"ConsumerDetailResp", string(data)}, " ")
}
