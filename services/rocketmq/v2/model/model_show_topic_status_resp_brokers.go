package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowTopicStatusRespBrokers struct {

	// 队列列表。
	Queues *[]ShowTopicStatusRespQueues `json:"queues,omitempty"`

	// 节点名称。
	BrokerName *string `json:"broker_name,omitempty"`
}

func (o ShowTopicStatusRespBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowTopicStatusRespBrokers struct{}"
	}

	return strings.Join([]string{"ShowTopicStatusRespBrokers", string(data)}, " ")
}
