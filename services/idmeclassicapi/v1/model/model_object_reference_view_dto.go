package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ObjectReferenceViewDto struct {

	// **参数解释：**  类名。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	Clazz *string `json:"clazz,omitempty"`

	// **参数解释：**  唯一标识。  **约束限制：**  不涉及。  **取值范围：**  -9223372036854775808到9223372036854775807的整数。  **默认取值：**  不涉及。
	Id string `json:"id"`
}

func (o ObjectReferenceViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ObjectReferenceViewDto struct{}"
	}

	return strings.Join([]string{"ObjectReferenceViewDto", string(data)}, " ")
}
