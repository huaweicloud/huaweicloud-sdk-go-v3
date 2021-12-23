package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 时间线
type StatisticsTimelineItem struct {
	// 键值

	Key *string `json:"key,omitempty"`
	// 对应键值的时间线

	Timeline *[]TimeLineItem `json:"timeline,omitempty"`
}

func (o StatisticsTimelineItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsTimelineItem struct{}"
	}

	return strings.Join([]string{"StatisticsTimelineItem", string(data)}, " ")
}
