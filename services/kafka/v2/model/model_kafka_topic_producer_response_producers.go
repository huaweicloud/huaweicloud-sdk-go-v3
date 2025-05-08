package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KafkaTopicProducerResponseProducers struct {

	// **参数解释**： 生产者地址。 **取值范围**： 不涉及
	ProducerAddress *string `json:"producer_address,omitempty"`

	// **参数解释**： broker地址。 **取值范围**： 不涉及
	BrokerAddress *string `json:"broker_address,omitempty"`

	// **参数解释**： 加入时间，以Unix时间戳显示。 **取值范围**： 不涉及
	JoinTime *int64 `json:"join_time,omitempty"`
}

func (o KafkaTopicProducerResponseProducers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaTopicProducerResponseProducers struct{}"
	}

	return strings.Join([]string{"KafkaTopicProducerResponseProducers", string(data)}, " ")
}
