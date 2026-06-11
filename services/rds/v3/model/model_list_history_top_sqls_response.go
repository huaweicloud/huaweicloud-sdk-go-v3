package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListHistoryTopSqlsResponse Response Object
type ListHistoryTopSqlsResponse struct {

	// 总记录数
	TotalCount *int32 `json:"total_count,omitempty"`

	// top sql 列表
	TopSqls        *[]TopSql `json:"top_sqls,omitempty"`
	HttpStatusCode int       `json:"-"`
}

func (o ListHistoryTopSqlsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListHistoryTopSqlsResponse struct{}"
	}

	return strings.Join([]string{"ListHistoryTopSqlsResponse", string(data)}, " ")
}
