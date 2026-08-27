package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListColdTableMetasResponse Response Object
type ListColdTableMetasResponse struct {

	// **参数解释**：  冷表元信息记录总数。  **取值范围**：  ≥0。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**：  冷表元记录列表。
	MetaInfo *[]ColdTableMetaInfo `json:"meta_info,omitempty"`

	// **参数解释**：  冷表数据量总大小（MB）。  **取值范围**：  ≥0。
	TotalDataSize  *float32 `json:"total_data_size,omitempty"`
	HttpStatusCode int      `json:"-"`
}

func (o ListColdTableMetasResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListColdTableMetasResponse struct{}"
	}

	return strings.Join([]string{"ListColdTableMetasResponse", string(data)}, " ")
}
