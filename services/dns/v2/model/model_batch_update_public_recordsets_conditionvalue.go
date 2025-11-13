package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdatePublicRecordsetsConditionvalue struct {

	// **参数解释：** 条件信息的键。 **约束限制：** 不涉及。 **取值范围：** - host_name：主机记录。  **默认取值：** 不涉及。
	Condition string `json:"condition"`

	// **参数解释：** 条件信息的值。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Value *string `json:"value,omitempty"`
}

func (o BatchUpdatePublicRecordsetsConditionvalue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePublicRecordsetsConditionvalue struct{}"
	}

	return strings.Join([]string{"BatchUpdatePublicRecordsetsConditionvalue", string(data)}, " ")
}
