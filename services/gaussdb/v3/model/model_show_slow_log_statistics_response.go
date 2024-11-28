package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogStatisticsResponse Response Object
type ShowSlowLogStatisticsResponse struct {

	// 慢日志统计信息列表。
	SlowLogList *[]ShowSlowLogStatisticsItem `json:"slow_log_list,omitempty"`

	// 总条数。
	TotalCount     *int32 `json:"total_count,omitempty"`
	HttpStatusCode int    `json:"-"`
}

func (o ShowSlowLogStatisticsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogStatisticsResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowLogStatisticsResponse", string(data)}, " ")
}
