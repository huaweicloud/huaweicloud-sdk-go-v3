package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CountResourceInstancesReqTags struct {

	// **参数解释**：标签键。 **约束限制**：不涉及。 **取值范围**：1~128个字符。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：标签值列表。 **约束限制**：最多10个标签值，每个标签值0~255个字符。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Values *[]string `json:"values,omitempty"`
}

func (o CountResourceInstancesReqTags) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourceInstancesReqTags struct{}"
	}

	return strings.Join([]string{"CountResourceInstancesReqTags", string(data)}, " ")
}
