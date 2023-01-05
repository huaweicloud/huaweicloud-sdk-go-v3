package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 带宽（BANDWIDTH）、入带宽（IN_BANDWIDTH）以及出带宽（OUT_BANDWIDTH）统计数据项
type BandwidthStatisticsTimelineItem struct {

	// 键值，包括带宽（BANDWIDTH）、入带宽（IN_BANDWIDTH）以及出带宽（OUT_BANDWIDTH）
	Key *string `json:"key,omitempty"`

	// 对应键值的时间线统计数据，包含两个字段，time字段值为时间点；num字段为time对应时间点与前一时间点间隔内的统计数值
	Timeline *[]TimeLineItem `json:"timeline,omitempty"`
}

func (o BandwidthStatisticsTimelineItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthStatisticsTimelineItem struct{}"
	}

	return strings.Join([]string{"BandwidthStatisticsTimelineItem", string(data)}, " ")
}
