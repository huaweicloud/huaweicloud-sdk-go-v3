package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KafkaTopicQuota topic流控配置
type KafkaTopicQuota struct {

	// topic名称
	Topic *string `json:"topic,omitempty"`

	// 生产者速率
	ProducerByteRate *int32 `json:"producer-byte-rate,omitempty"`

	// 消费者速率
	ConsumerByteRate *int32 `json:"consumer-byte-rate,omitempty"`
}

func (o KafkaTopicQuota) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaTopicQuota struct{}"
	}

	return strings.Join([]string{"KafkaTopicQuota", string(data)}, " ")
}
