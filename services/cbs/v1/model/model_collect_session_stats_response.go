package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type CollectSessionStatsResponse struct {

	// 统计周期目前支持year、month、week、day。
	Interval *string `json:"interval,omitempty" xml:"interval"`

	// 所在时区，默认为\"utc\"。例如：中国东八区为\"+08:00\"；美国西五区为\"-05:00\"。
	TimeZone *string `json:"time_zone,omitempty" xml:"time_zone"`

	Total *SessionStatsTotal `json:"total,omitempty" xml:"total"`

	// 会话间隔统计数据。
	Intervals *[]SessionStatsIntervals `json:"intervals,omitempty" xml:"intervals"`

	// 统计开始的utc时间。
	Startutc *int64 `json:"startutc,omitempty" xml:"startutc"`

	// 统计结束的utc时间。
	Endutc         *int64 `json:"endutc,omitempty" xml:"endutc"`
	HttpStatusCode int    `json:"-"`
}

func (o CollectSessionStatsResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CollectSessionStatsResponse struct{}"
	}

	return strings.Join([]string{"CollectSessionStatsResponse", string(data)}, " ")
}
