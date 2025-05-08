package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Brokers struct {

	// **参数解释**： Topic关联代理名称。 **取值范围**： 不涉及。
	BrokerName *string `json:"broker_name,omitempty"`

	// **参数解释**： 关联代理的队列详情。
	Queues *[]Queue `json:"queues,omitempty"`
}

func (o Brokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Brokers struct{}"
	}

	return strings.Join([]string{"Brokers", string(data)}, " ")
}
