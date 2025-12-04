package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowKafkaTopicDetailResponseBody struct {

	// **参数解释**： Topic名称。  **取值范围**： 不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 创建时间。  **取值范围**： 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 分区列表。
	Partitions *[]ShowKafkaTopicDetailResponsePartitions `json:"partitions,omitempty"`
}

func (o ShowKafkaTopicDetailResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKafkaTopicDetailResponseBody struct{}"
	}

	return strings.Join([]string{"ShowKafkaTopicDetailResponseBody", string(data)}, " ")
}
