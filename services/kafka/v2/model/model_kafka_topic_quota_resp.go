package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KafkaTopicQuotaResp Topic流控配置
type KafkaTopicQuotaResp struct {

	// Topic名称
	Topic *string `json:"topic,omitempty"`

	// 生产者速率
	ProducerByteRate *int32 `json:"producer-byte-rate,omitempty"`

	// 消费者速率
	ConsumerByteRate *int32 `json:"consumer-byte-rate,omitempty"`
}

func (o KafkaTopicQuotaResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaTopicQuotaResp struct{}"
	}

	return strings.Join([]string{"KafkaTopicQuotaResp", string(data)}, " ")
}
