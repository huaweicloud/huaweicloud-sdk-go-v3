package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListSqlStatisticsResponse Response Object
type ListSqlStatisticsResponse struct {

	// 最新重置时间
	ResetViewLastTime *int64 `json:"reset_view_last_time,omitempty"`

	// 总的个数
	TotalCount *int32 `json:"total_count,omitempty"`

	// 数据列表
	List           *[]Statistic `json:"list,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ListSqlStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListSqlStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ListSqlStatisticsResponse", string(data)}, " ")
}
