package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 不同键值对应的时间线
type BandwidthStatisticsTimelineItem struct {

	// 键值，包括带宽（BANDWIDTH）、入带宽（IN_BANDWIDTH）以及出带宽（OUT_BANDWIDTH）
	Key *string `json:"key,omitempty" xml:"key"`

	// 对应键值的时间线统计数据
	Timeline *[]TimeLineItem `json:"timeline,omitempty" xml:"timeline"`
}

func (o BandwidthStatisticsTimelineItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BandwidthStatisticsTimelineItem struct{}"
	}

	return strings.Join([]string{"BandwidthStatisticsTimelineItem", string(data)}, " ")
}
