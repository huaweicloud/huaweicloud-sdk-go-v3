package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListTimingOffConfigInfoResponseInfo 定时关闭防护功能配置详情
type ListTimingOffConfigInfoResponseInfo struct {

	// **参数解释**: 自动关闭防护周期列表。1代表周一；2代表周二；3代表周三；4代表周四；5代表周五；6代表周六；7代表周日。 **取值范围**: 最少0条，最多7条
	WeekOffList *[]int32 `json:"week_off_list,omitempty"`

	// **参数解释**: 自动关闭防护时间段 **取值范围**: 最少0条，最多5条
	TimingRangeList *[]TimingRangeConfigInfo `json:"timing_range_list,omitempty"`
}

func (o ListTimingOffConfigInfoResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListTimingOffConfigInfoResponseInfo struct{}"
	}

	return strings.Join([]string{"ListTimingOffConfigInfoResponseInfo", string(data)}, " ")
}
