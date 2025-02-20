package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ExportFullSqlDetailsResponse Response Object
type ExportFullSqlDetailsResponse struct {

	// 全量SQL明细列表。
	FullSqlDetails *[]FullSqlDetail `json:"full_sql_details,omitempty"`

	// 总数。
	Total          *int32 `json:"total,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ExportFullSqlDetailsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ExportFullSqlDetailsResponse struct{}"
	}

	return strings.Join([]string{"ExportFullSqlDetailsResponse", string(data)}, " ")
}
