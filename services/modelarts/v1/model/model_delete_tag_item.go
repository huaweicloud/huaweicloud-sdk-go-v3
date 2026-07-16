package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteTagItem **参数解释**：资源标签信息。
type DeleteTagItem struct {

	// **参数解释**：标签的key。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Key string `json:"key"`

	// **参数解释**：标签的value。 **约束限制**：不涉及。 **取值范围**：不涉及。 **默认取值**：不涉及。
	Value *string `json:"value,omitempty"`
}

func (o DeleteTagItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteTagItem struct{}"
	}

	return strings.Join([]string{"DeleteTagItem", string(data)}, " ")
}
