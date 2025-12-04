package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateInstanceTopicReqTopics struct {

	// **参数解释**： Topic名称，不支持修改。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 老化时间，单位小时。 **约束限制**： 不涉及。 **取值范围**： 0-720。 **默认取值**： 不涉及。
	RetentionTime *int32 `json:"retention_time,omitempty"`

	// **参数解释**： 是否同步复制。 **约束限制**： 不涉及。 **取值范围**： - true：开启同步复制。 - false：不开启同步复制。 **默认取值**： 不涉及。
	SyncReplication *bool `json:"sync_replication,omitempty"`

	// **参数解释**： 是否同步落盘。 **约束限制**： 不涉及。 **取值范围**： - true：开启同步落盘。 - false：不开启同步落盘。 **默认取值**： 不涉及。
	SyncMessageFlush *bool `json:"sync_message_flush,omitempty"`

	// **参数解释**： 分区数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NewPartitionNumbers *int32 `json:"new_partition_numbers,omitempty"`

	// **参数解释**： 增加分区时指定broker列表。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	NewPartitionBrokers *[]int32 `json:"new_partition_brokers,omitempty"`

	// **参数解释**： Topic配置。 **约束限制**： 不涉及。
	TopicOtherConfigs *[]UpdateInstanceTopicReqTopicOtherConfigs `json:"topic_other_configs,omitempty"`

	// **参数解释**： Topic描述。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TopicDesc *string `json:"topic_desc,omitempty"`
}

func (o UpdateInstanceTopicReqTopics) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateInstanceTopicReqTopics struct{}"
	}

	return strings.Join([]string{"UpdateInstanceTopicReqTopics", string(data)}, " ")
}
