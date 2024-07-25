package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SortInfoVo struct {
	CharacterSet *CharacterSetEnum `json:"characterSet,omitempty"`

	// **参数解释：**  按某个字段进行排序。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	OrderBy *string `json:"orderBy,omitempty"`

	// **参数解释：**  排序方向。  **约束限制：**  不涉及。  **取值范围：**  - ASC：表示升序。 - DESC：表示降序。  **默认取值：**  ASC。
	Sort *string `json:"sort,omitempty"`

	// **参数解释：**  排序信息。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	SortInfo *string `json:"sortInfo,omitempty"`

	// **参数解释：**  排序信息字段。  **约束限制：**  不涉及。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	SortInfoOrderBy *string `json:"sortInfoOrderBy,omitempty"`
}

func (o SortInfoVo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SortInfoVo struct{}"
	}

	return strings.Join([]string{"SortInfoVo", string(data)}, " ")
}
