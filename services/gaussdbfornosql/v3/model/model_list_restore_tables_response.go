package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRestoreTablesResponse Response Object
type ListRestoreTablesResponse struct {

	// 总记录数。
	TotalCount *int32 `json:"total_count,omitempty"`

	// 数据表名称列表。
	TableNames     *[]string `json:"table_names,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListRestoreTablesResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRestoreTablesResponse struct{}"
	}

	return strings.Join([]string{"ListRestoreTablesResponse", string(data)}, " ")
}
