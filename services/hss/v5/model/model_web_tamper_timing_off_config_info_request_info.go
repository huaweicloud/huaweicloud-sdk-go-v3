package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type WebTamperTimingOffConfigInfoRequestInfo struct {

	// 关闭防护周期
	WeekOffList *[]int32 `json:"week_off_list,omitempty"`

	// 时间段
	TimingRangeList *[]TimingRangeConfigInfo `json:"timing_range_list,omitempty"`
}

func (o WebTamperTimingOffConfigInfoRequestInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WebTamperTimingOffConfigInfoRequestInfo struct{}"
	}

	return strings.Join([]string{"WebTamperTimingOffConfigInfoRequestInfo", string(data)}, " ")
}
