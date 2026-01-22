package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateBindingBody struct {

	// 要投递的Exchange或Queue名称
	Destination string `json:"destination"`

	// **参数解释**： 绑定键值，用于告知Exchange应该将消息投递到哪些Queue或Exchange中。 **约束限制**： - 不允许包含+和~。 - 最长128个字符。 - 不能包含中文。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	RoutingKey string `json:"routing_key"`

	// **参数解释**： 绑定目标端类型。 **约束限制**： [不涉及。](tag:sbc,cmcc,tm,hk_tm,ax,hk_sbc)[AMQP版本只支持绑定Queue。](tag:hws,hws_hk,hws_eu) **取值范围**： - exchange：交换机。 - queue：队列。 **默认取值**： 不涉及。
	DestinationType string `json:"destination_type"`
}

func (o CreateBindingBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBindingBody struct{}"
	}

	return strings.Join([]string{"CreateBindingBody", string(data)}, " ")
}
