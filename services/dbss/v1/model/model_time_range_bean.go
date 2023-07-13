package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TimeRangeBean struct {

	// 开始时间，必须和end_time成对出现。格式必须为yyyy-MM-dd HH:mm:ss。UTC时间
	EndTime *string `json:"end_time,omitempty"`

	// 结束时间，必须和start_time成对出现。格式必须为yyyy-MM-dd HH:mm:ss。UTC时间
	StartTime *string `json:"start_time,omitempty"`

	// 请求查询的时间段，值为： • HALF_HOUR • HOUR • THREE_HOUR • TWELVE_HOUR • DAY（24小时） • WEEK（7天） • MONTH（30天） 枚举值：  HALF_HOUR  HOUR  THREE_HOUR  TWELVE_HOUR  DAY  WEEK  MONTH
	TimeRange *string `json:"time_range,omitempty"`
}

func (o TimeRangeBean) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TimeRangeBean struct{}"
	}

	return strings.Join([]string{"TimeRangeBean", string(data)}, " ")
}
