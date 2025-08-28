package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTimingOffConfigInfoResponseInfo 定时关闭防护功能配置详情
type ListTimingOffConfigInfoResponseInfo struct {

	// **参数解释**: 自动关闭防护周期列表 **取值范围**: 最少0条，最多7条
	WeekOffList *[]int32 `json:"week_off_list,omitempty"`

	// **参数解释**: 自动关闭防护时间段 **取值范围**: 最少0条，最多5条
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
