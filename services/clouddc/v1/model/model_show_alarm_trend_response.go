package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowAlarmTrendResponse Response Object
type ShowAlarmTrendResponse struct {

	// 故障服务器数趋势
	FaultHosts *[]FaultHostInfo `json:"fault_hosts,omitempty"`

	// 告警每日新增趋势
	AlarmDailyNewTrends *[]AlarmDailyTrend `json:"alarm_daily_new_trends,omitempty"`
	HttpStatusCode      int                `json:"-"`
}

func (o ShowAlarmTrendResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowAlarmTrendResponse struct{}"
	}

	return strings.Join([]string{"ShowAlarmTrendResponse", string(data)}, " ")
}
