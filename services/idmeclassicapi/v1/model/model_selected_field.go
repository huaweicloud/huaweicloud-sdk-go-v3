package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SelectedField struct {

	// **参数解释：**  字段名称，如果是子参考对象的属性，则为“参考对象.属性名称”，例如：“master.name”。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Name string `json:"name"`

	// **参数解释：**  字段别名。如果不填，默认使用name参数的值。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	NameAs *string `json:"nameAs,omitempty"`
}

func (o SelectedField) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SelectedField struct{}"
	}

	return strings.Join([]string{"SelectedField", string(data)}, " ")
}
