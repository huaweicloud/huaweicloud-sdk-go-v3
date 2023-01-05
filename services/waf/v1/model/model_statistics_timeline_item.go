package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 不同键值对应的时间线
type StatisticsTimelineItem struct {

	// 键值，包括请求总量（ACCESS）、Bot攻击防护（CRAWLER）、攻击总量（ATTACK）、Web基础防护（WEB_ATTACK）、精准防护（PRECISE）、CC攻击防护（CC）
	Key *string `json:"key,omitempty"`

	// 对应键值的时间线统计数据
	Timeline *[]TimeLineItem `json:"timeline,omitempty"`
}

func (o StatisticsTimelineItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StatisticsTimelineItem struct{}"
	}

	return strings.Join([]string{"StatisticsTimelineItem", string(data)}, " ")
}
