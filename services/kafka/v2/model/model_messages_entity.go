package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MessagesEntity struct {

	// **参数解释**： Topic名称。 **取值范围**： 不涉及。
	Topic *string `json:"topic,omitempty"`

	// **参数解释**： 消息所在的分区。 **取值范围**： 不涉及。
	Partition *int32 `json:"partition,omitempty"`

	// **参数解释**： 消息key。 **取值范围**： 不涉及。
	Key *string `json:"key,omitempty"`

	// **参数解释**： 消息内容。 **取值范围**： 不涉及。
	Value *string `json:"value,omitempty"`

	// **参数解释**： 消息大小。 **取值范围**： 不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释**： 生产消息的时间。 格式为Unix时间戳。单位为毫秒。 **取值范围**： 不涉及。
	Timestamp *int64 `json:"timestamp,omitempty"`

	// **参数解释**： 大数据标识。 **取值范围**： 不涉及。
	HugeMessage *bool `json:"huge_message,omitempty"`

	// **参数解释**： 消息偏移量。 **取值范围**： 不涉及。
	MessageOffset *int64 `json:"message_offset,omitempty"`

	// **参数解释**： 消息ID。 **取值范围**： 不涉及。
	MessageId *string `json:"message_id,omitempty"`

	// **参数解释**： 应用ID。 **取值范围**： 不涉及。
	AppId *string `json:"app_id,omitempty"`

	// **参数解释**： 消息标签。 **取值范围**： 不涉及。
	Tag *string `json:"tag,omitempty"`
}

func (o MessagesEntity) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MessagesEntity struct{}"
	}

	return strings.Join([]string{"MessagesEntity", string(data)}, " ")
}
