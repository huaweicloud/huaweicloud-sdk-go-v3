package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TopicEntity struct {

	// **参数解释**： 是否为默认策略。 **取值范围**： - true：默认策略。 - false：不是默认策略。
	PoliciesOnly *bool `json:"policiesOnly,omitempty"`

	// **参数解释**： Topic名称。 **取值范围**： 不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**： 副本数，配置数据的可靠性。 **取值范围**： 不涉及
	Replication *int32 `json:"replication,omitempty"`

	// **参数解释**： Topic分区数，设置消费的并发数。 **取值范围**： 不涉及
	Partition *int32 `json:"partition,omitempty"`

	// **参数解释**： 消息老化时间。 **取值范围**： 0-720
	RetentionTime *int32 `json:"retention_time,omitempty"`

	// **参数解释**： 是否开启同步复制，默认关闭。 **取值范围**： - true：开启，客户端生产消息时相应的也要设置acks=-1，否则不生效。 - false：关闭。
	SyncReplication *bool `json:"sync_replication,omitempty"`

	// **参数解释**： 是否使用同步落盘。默认值为false。同步落盘会导致性能降低。 **取值范围**： - true：同步落盘。 - false：不同步落盘。
	SyncMessageFlush *bool `json:"sync_message_flush,omitempty"`

	// **参数解释**： 扩展配置。
	ExternalConfigs *interface{} `json:"external_configs,omitempty"`

	// **参数解释**： Topic类型。 **取值范围**： - 0：普通Topic。 - 1：系统(内部)Topic。
	TopicType *int32 `json:"topic_type,omitempty"`

	// **参数解释**： Topic其他配置。
	TopicOtherConfigs *[]TopicEntityTopicOtherConfigs `json:"topic_other_configs,omitempty"`

	// **参数解释**： Topic描述。 **取值范围**： 不涉及
	TopicDesc *string `json:"topic_desc,omitempty"`

	// **参数解释**： Topic创建时间。 **取值范围**： 不涉及
	CreatedAt *int64 `json:"created_at,omitempty"`
}

func (o TopicEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TopicEntity struct{}"
	}

	return strings.Join([]string{"TopicEntity", string(data)}, " ")
}
