package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CreateInstanceTopicReq struct {

	// **参数解释**： Topic名称。 **约束限制**： 长度为3-200，以字母开头且只支持大小写字母、中横线、下划线、点以及数字。 **取值范围**： 不涉及 **默认取值**： 不涉及。
	Id string `json:"id"`

	// **参数解释**： 副本数，配置数据的可靠性。副本数和代理数有关，如果有3个代理，最大副本数是3。 **约束限制**： 不涉及。 **取值范围**： 不涉及 **默认取值**： 不涉及。
	Replication *int32 `json:"replication,omitempty"`

	// **参数解释**： 是否使用同步落盘，同步落盘会导致性能降低。 **约束限制**： 不涉及。 **取值范围**： - true：同步落盘。 - false：不使用同步落盘。 **默认取值**： false
	SyncMessageFlush *bool `json:"sync_message_flush,omitempty"`

	// **参数解释**： Topic分区数，设置消费的并发数。 **约束限制**： 不涉及。 **取值范围**： 1-200。 **默认取值**： 不涉及。
	Partition *int32 `json:"partition,omitempty"`

	// **参数解释**： 是否开启同步复制。 **约束限制**： 不涉及。 **取值范围**： - true：开启。开启后，客户端生产消息时相应的也要设置acks=-1，否则不生效。 - false：不开启。 **默认取值**： false。
	SyncReplication *bool `json:"sync_replication,omitempty"`

	// **参数解释**： 消息老化时间。 **约束限制**： 不涉及。 **取值范围**： 1-720 **默认取值**： 72
	RetentionTime *int32 `json:"retention_time,omitempty"`

	// **参数解释**： Topic配置
	TopicOtherConfigs *[]CreateInstanceTopicReqTopicOtherConfigs `json:"topic_other_configs,omitempty"`

	// **参数解释**： Topic描述。 **约束限制**： 不涉及。 **取值范围**： 0-200个字符。 **默认取值**： 不涉及。
	TopicDesc *string `json:"topic_desc,omitempty"`
}

func (o CreateInstanceTopicReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateInstanceTopicReq struct{}"
	}

	return strings.Join([]string{"CreateInstanceTopicReq", string(data)}, " ")
}
