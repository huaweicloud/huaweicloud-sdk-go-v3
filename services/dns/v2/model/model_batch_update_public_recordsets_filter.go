package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdatePublicRecordsetsFilter struct {

	// **参数解释：** 过滤条件之间的关系。 **约束限制：** 不涉及。 **取值范围：** - OR：或 - AND：与  **默认取值：** 不涉及。
	Relation string `json:"relation"`

	// **参数解释：** 条件列表。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Conditions []BatchUpdatePublicRecordsetsConditionvalue `json:"conditions"`
}

func (o BatchUpdatePublicRecordsetsFilter) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePublicRecordsetsFilter struct{}"
	}

	return strings.Join([]string{"BatchUpdatePublicRecordsetsFilter", string(data)}, " ")
}
