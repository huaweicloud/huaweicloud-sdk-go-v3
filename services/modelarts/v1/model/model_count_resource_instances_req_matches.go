package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CountResourceInstancesReqMatches struct {

	// **参数解释**：匹配字段名。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：匹配值。 **约束限制**：不涉及。 **取值范围**：1~255个字符。 **默认取值**：不涉及。
	Value string `json:"value"`
}

func (o CountResourceInstancesReqMatches) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CountResourceInstancesReqMatches struct{}"
	}

	return strings.Join([]string{"CountResourceInstancesReqMatches", string(data)}, " ")
}
