package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type Brokers struct {

	// Topic关联代理名称。
	BrokerName *string `json:"broker_name,omitempty"`

	// 关联代理的队列详情。
	Queues *[]Queue `json:"queues,omitempty"`
}

func (o Brokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Brokers struct{}"
	}

	return strings.Join([]string{"Brokers", string(data)}, " ")
}
