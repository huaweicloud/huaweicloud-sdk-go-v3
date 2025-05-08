package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPartitionMessageEntity 消息体。
type ShowPartitionMessageEntity struct {

	// **参数解释**： 消息的key。 **取值范围**： 不涉及
	Key *string `json:"key,omitempty"`

	// **参数解释**： 消息内容。 **取值范围**： 不涉及
	Value *string `json:"value,omitempty"`

	// **参数解释**： Topic名称。 **取值范围**： 不涉及
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 分区编号。 **取值范围**： 不涉及
	Partition *int32 `json:"partition,omitempty"`

	// **参数解释**： 消息位置。 **取值范围**： 不涉及
	MessageOffset *int64 `json:"message_offset,omitempty"`

	// **参数解释**： 消息大小，单位字节。 **取值范围**： 不涉及
	Size *int32 `json:"size,omitempty"`

	// **参数解释**： 生产消息的时间。 格式为Unix时间戳。单位为毫秒。 **取值范围**： 不涉及
	Timestamp *int64 `json:"timestamp,omitempty"`
}

func (o ShowPartitionMessageEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPartitionMessageEntity struct{}"
	}

	return strings.Join([]string{"ShowPartitionMessageEntity", string(data)}, " ")
}
