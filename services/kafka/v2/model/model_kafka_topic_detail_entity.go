package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KafkaTopicDetailEntity struct {

	// **参数解释**： Topic名称。  **取值范围**： 不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 创建时间。  **取值范围**： 不涉及。
	CreateTime *int64 `json:"create_time,omitempty"`

	// **参数解释**： 分区列表。
	Partitions *[]KafkaTopicDetailEntityPartitions `json:"partitions,omitempty"`
}

func (o KafkaTopicDetailEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaTopicDetailEntity struct{}"
	}

	return strings.Join([]string{"KafkaTopicDetailEntity", string(data)}, " ")
}
