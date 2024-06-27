package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AlarmStatisticsQuery struct {

	// 结束时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 开始时间
	StartTime *int64 `json:"start_time,omitempty"`
}

func (o AlarmStatisticsQuery) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmStatisticsQuery struct{}"
	}

	return strings.Join([]string{"AlarmStatisticsQuery", string(data)}, " ")
}
