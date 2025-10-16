package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSessionTopSqlStatisticsResponse Response Object
type ListSessionTopSqlStatisticsResponse struct {

	// **参数解释**: 统计总条数。 **取值范围**: 不涉及。
	Total *int32 `json:"total,omitempty"`

	// **参数解释**: 统计数据。
	TopSqlInfo     *[]SessionTopSqlStatisticInfo `json:"top_sql_info,omitempty"`
	HttpStatusCode int                           `json:"-"`
}

func (o ListSessionTopSqlStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSessionTopSqlStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListSessionTopSqlStatisticsResponse", string(data)}, " ")
}
