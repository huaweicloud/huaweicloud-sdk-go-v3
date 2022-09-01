package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopicBrokers struct {

	// 代理名称。
	BrokerName *string `json:"broker_name,omitempty" xml:"broker_name"`

	// 读队列个数。
	ReadQueueNum float32 `json:"read_queue_num,omitempty" xml:"read_queue_num"`

	// 写队列个数。
	WriteQueueNum float32 `json:"write_queue_num,omitempty" xml:"write_queue_num"`
}

func (o TopicBrokers) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicBrokers struct{}"
	}

	return strings.Join([]string{"TopicBrokers", string(data)}, " ")
}
