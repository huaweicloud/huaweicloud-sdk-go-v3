package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// MigrationRocketMqTopicConfig **参数解释**： RocketMQ元数据迁移，RocketMQ Topic元数据。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
type MigrationRocketMqTopicConfig struct {

	// **参数解释**： Topic名称。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 不涉及。
	TopicName *string `json:"topic_name,omitempty"`

	// **参数解释**： 是否有序消息。 **约束限制**： 不涉及。 **取值范围**： - true：有序消息。 - false：无序消息。 **默认取值**： false。
	Order *bool `json:"order,omitempty"`

	// **参数解释**： Topic权限。 **约束限制**： 不涉及。 **取值范围**： - PUB：拥有发布权限。 - SUB：拥有订阅权限。 - PUB|SUB：拥有发布订阅权限。 - DENY：无权限。 **默认取值**： 6。
	Perm *int32 `json:"perm,omitempty"`

	// **参数解释**： 读队列个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 16。
	ReadQueueNums *int32 `json:"read_queue_nums,omitempty"`

	// **参数解释**： 写队列个数。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 16。
	WriteQueueNums *int32 `json:"write_queue_nums,omitempty"`

	// **参数解释**： Topic过滤类型。 **约束限制**： 不涉及。 **取值范围**： - SINGLE_TAG：单标签   - MULTI_TAG：多标签 **默认取值**： 不涉及。
	TopicFilterType *string `json:"topic_filter_type,omitempty"`

	// **参数解释**： Topic系统标志位。 **约束限制**： 不涉及。 **取值范围**： 不涉及。 **默认取值**： 0。
	TopicSysFlag *int32 `json:"topic_sys_flag,omitempty"`
}

func (o MigrationRocketMqTopicConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MigrationRocketMqTopicConfig struct{}"
	}

	return strings.Join([]string{"MigrationRocketMqTopicConfig", string(data)}, " ")
}
