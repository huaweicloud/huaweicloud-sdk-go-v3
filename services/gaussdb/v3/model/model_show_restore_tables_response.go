package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowRestoreTablesResponse Response Object
type ShowRestoreTablesResponse struct {

	// 数据库总数。
	TotalDatabases *int32 `json:"total_databases,omitempty"`

	// 数据库信息。
	Databases      *[]RestoreDatabaseInfos `json:"databases,omitempty"`
	HttpStatusCode int                     `json:"-"`
}

func (o ShowRestoreTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowRestoreTablesResponse struct{}"
	}

	return strings.Join([]string{"ShowRestoreTablesResponse", string(data)}, " ")
}
