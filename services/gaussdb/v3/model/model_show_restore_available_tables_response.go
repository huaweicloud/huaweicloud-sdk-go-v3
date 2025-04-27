package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestoreAvailableTablesResponse Response Object
type ShowRestoreAvailableTablesResponse struct {

	// **参数解释**:  实例总表数。  **取值范围**：  不涉及。
	TotalTables *int32 `json:"total_tables,omitempty"`

	// **参数解释**：  数据库信息。
	Databases      *[]RestoreDatabaseInfos `json:"databases,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowRestoreAvailableTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestoreAvailableTablesResponse struct{}"
	}

	return strings.Join([]string{"ShowRestoreAvailableTablesResponse", string(data)}, " ")
}
