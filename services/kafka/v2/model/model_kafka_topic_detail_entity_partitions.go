package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type KafkaTopicDetailEntityPartitions struct {

	// **参数解释**： 分区编号。  **取值范围**： 不涉及。
	Partition *int32 `json:"partition,omitempty"`

	// **参数解释**： leader副本所在节点的ID。  **取值范围**： 不涉及。
	Leader *int32 `json:"leader,omitempty"`

	// **参数解释**： 副本列表。
	Replicas *[]int32 `json:"replicas,omitempty"`
}

func (o KafkaTopicDetailEntityPartitions) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KafkaTopicDetailEntityPartitions struct{}"
	}

	return strings.Join([]string{"KafkaTopicDetailEntityPartitions", string(data)}, " ")
}
