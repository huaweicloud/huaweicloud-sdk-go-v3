package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ListTimingOffConfigInfoResponseInfo struct {

	// 关闭防护周期
	WeekOffList *[]int32 `json:"week_off_list,omitempty"`

	// 时间段
	TimingRangeList *[]TimingRangeConfigInfo `json:"timing_range_list,omitempty"`

	// **参数解释**: 总数 **取值范围**: 最小值0，最大值2147483647
	TotalNum *int32 `json:"total_num,omitempty"`
}

func (o ListTimingOffConfigInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTimingOffConfigInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"ListTimingOffConfigInfoResponseInfo", string(data)}, " ")
}
