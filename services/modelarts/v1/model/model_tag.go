package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Tag **参数解释**：资源标签的详细信息。
type Tag struct {

	// **参数解释**：标签的key。 **取值范围**：不涉及。
	Key string `json:"key"`

	// **参数解释**：标签的value。 **取值范围**：不涉及。
	Value string `json:"value"`
}

func (o Tag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Tag struct{}"
	}

	return strings.Join([]string{"Tag", string(data)}, " ")
}
