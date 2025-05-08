package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopicBrokers struct {

	// **参数解释**： 代理名称。 **取值范围**： 不涉及。
	BrokerName *string `json:"broker_name,omitempty"`

	// **参数解释**： 读队列个数。 **取值范围**： 不涉及。
	ReadQueueNum float32 `json:"read_queue_num,omitempty"`

	// **参数解释**： 写队列个数。 **取值范围**： 不涉及。
	WriteQueueNum float32 `json:"write_queue_num,omitempty"`
}

func (o TopicBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicBrokers struct{}"
	}

	return strings.Join([]string{"TopicBrokers", string(data)}, " ")
}
