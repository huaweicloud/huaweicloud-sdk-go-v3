package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowSlowLogStatisticsNewResponse Response Object
type ShowSlowLogStatisticsNewResponse struct {

	// 慢日志统计列表
	StatisticsList *[]SlowLogStatistics `json:"statistics_list,omitempty"`
	HttpStatusCode int                  `json:"-"`
}

func (o ShowSlowLogStatisticsNewResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowSlowLogStatisticsNewResponse struct{}"
	}

	return strings.Join([]string{"ShowSlowLogStatisticsNewResponse", string(data)}, " ")
}
