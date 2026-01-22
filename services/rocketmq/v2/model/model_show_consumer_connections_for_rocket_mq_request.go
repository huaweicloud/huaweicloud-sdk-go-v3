package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowConsumerConnectionsForRocketMqRequest Request Object
type ShowConsumerConnectionsForRocketMqRequest struct {

	// **参数解释**： 实例ID。获取方法如下：调用“查询所有实例列表”接口，从响应体中获取实例ID。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	InstanceId string `json:"instance_id"`

	// **参数解释**： 消费组名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Group string `json:"group"`

	// **参数解释**： 查询数量。 **约束限制**： 不涉及。 **取值范围**： 1~50。 **默认取值**： 不涉及。
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**： 偏移量，表示从此偏移量开始查询， offset大于等于0。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Offset *int32 `json:"offset,omitempty"`

	// **参数解释**： 是否查询消费者详细列表，参数为“true”则表示查询详细列表，否则表示查询简易列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	IsDetail *bool `json:"is_detail,omitempty"`
}

func (o ShowConsumerConnectionsForRocketMqRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowConsumerConnectionsForRocketMqRequest struct{}"
	}

	return strings.Join([]string{"ShowConsumerConnectionsForRocketMqRequest", string(data)}, " ")
}
