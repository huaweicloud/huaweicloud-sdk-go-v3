package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ListWeeklyReportsResponse struct {

	// 一周内DDoS拦截次数
	DdosInterceptTimes *int32 `json:"ddos_intercept_times,omitempty"`

	// 一周的攻击次数统计数据
	Weekdata *[]WeeklyCount `json:"weekdata,omitempty"`

	// 被攻击次数排名前10的IP地址
	Top10          *[]WeeklyTop10 `json:"top10,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ListWeeklyReportsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListWeeklyReportsResponse struct{}"
	}

	return strings.Join([]string{"ListWeeklyReportsResponse", string(data)}, " ")
}
