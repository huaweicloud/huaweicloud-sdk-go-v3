package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListArchivableTablesResponse Response Object
type ListArchivableTablesResponse struct {

	// **参数解释**：  总记录数。  **取值范围**：  大于等于0。
	TotalCount *int32 `json:"total_count,omitempty"`

	// **参数解释**：  可归档表信息。
	ArchivableTables *[]ArchivableTable `json:"archivable_tables,omitempty"`
	HttpStatusCode   int                `json:"-"`
}

func (o ListArchivableTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListArchivableTablesResponse struct{}"
	}

	return strings.Join([]string{"ListArchivableTablesResponse", string(data)}, " ")
}
