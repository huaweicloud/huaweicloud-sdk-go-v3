package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BatchUpdatePublicRecordsetsUpdateValue struct {

	// **参数解释：** 待修改信息的键。 **约束限制：** 不涉及。 **取值范围：** - host_name：主机记录。 - rr_value: 记录值。  **默认取值：** 不涉及。
	Key string `json:"key"`

	// **参数解释：** 待修改信息的值。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Values []string `json:"values"`
}

func (o BatchUpdatePublicRecordsetsUpdateValue) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchUpdatePublicRecordsetsUpdateValue struct{}"
	}

	return strings.Join([]string{"BatchUpdatePublicRecordsetsUpdateValue", string(data)}, " ")
}
