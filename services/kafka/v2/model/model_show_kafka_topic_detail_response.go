package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKafkaTopicDetailResponse Response Object
type ShowKafkaTopicDetailResponse struct {

	// 查询Kafka主题详情
	Body           *[]KafkaTopicDetailEntity `json:"body,omitempty"`
	HttpStatusCode int                       `json:"-"`
}

func (o ShowKafkaTopicDetailResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaTopicDetailResponse struct{}"
	}

	return strings.Join([]string{"ShowKafkaTopicDetailResponse", string(data)}, " ")
}
