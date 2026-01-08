package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTag **参数解释**：ELB资源的标签列表。
type ListTag struct {

	// **参数解释**：标签的键。  **取值范围**：最大长度36个unicode字符。
	Key string `json:"key"`

	// **参数解释**：标签的值列表。  **取值范围**：列表中的每个值最大长度43个unicode字符，可以为空字符串。
	Values []string `json:"values"`
}

func (o ListTag) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTag struct{}"
	}

	return strings.Join([]string{"ListTag", string(data)}, " ")
}
