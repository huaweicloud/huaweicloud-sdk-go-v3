package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PageInfoViewDto struct {

	// **参数解释：**  当前页。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	CurPage *int32 `json:"curPage,omitempty"`

	// **参数解释：**  每页大小。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	PageSize *int32 `json:"pageSize,omitempty"`

	// **参数解释：**  总行数。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TotalRows *int32 `json:"totalRows,omitempty"`

	// **参数解释：**  总页数。  **取值范围：**  不涉及。  **默认取值：**  不涉及。
	TotalPages *int32 `json:"totalPages,omitempty"`
}

func (o PageInfoViewDto) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PageInfoViewDto struct{}"
	}

	return strings.Join([]string{"PageInfoViewDto", string(data)}, " ")
}
