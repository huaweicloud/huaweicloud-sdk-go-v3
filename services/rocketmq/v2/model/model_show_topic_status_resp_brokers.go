package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowTopicStatusRespBrokers struct {

	// **参数解释**： 队列列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Queues *[]ShowTopicStatusRespQueues `json:"queues,omitempty"`

	// **参数解释**： 节点名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	BrokerName *string `json:"broker_name,omitempty"`
}

func (o ShowTopicStatusRespBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicStatusRespBrokers struct{}"
	}

	return strings.Join([]string{"ShowTopicStatusRespBrokers", string(data)}, " ")
}
