package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CreateQueueResponse struct {

	// 队列ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 队列的名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 仅Kafka队列才有该响应参数。  使用Kafka SDK时的Kafka topic的ID。
	KafkaTopic     *string `json:"kafka_topic,omitempty" xml:"kafka_topic"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateQueueResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateQueueResponse struct{}"
	}

	return strings.Join([]string{"CreateQueueResponse", string(data)}, " ")
}
