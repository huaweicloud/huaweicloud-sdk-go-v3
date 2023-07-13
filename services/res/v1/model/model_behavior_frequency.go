package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BehaviorFrequency
type BehaviorFrequency struct {

	// 行为类型： - view，曝光 - click，点击 - collect，收藏 - uncollect，取消收藏 - search_click，搜索后点击 - comment，评论 - share，分享 - like，点赞 - dislike，点衰 - grade，评分 - consume，消费 - use，观看视频/听音乐/阅读 - download，下载 - tip，打赏 - subscribe，关注
	BehaviorType string `json:"behavior_type"`

	// 最小次数。
	LowerLimit *int32 `json:"lower_limit,omitempty"`

	// 最大次数。
	UpperLimit *int32 `json:"upper_limit,omitempty"`

	// 时间区间。
	TimeInterval int32 `json:"time_interval"`
}

func (o BehaviorFrequency) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BehaviorFrequency struct{}"
	}

	return strings.Join([]string{"BehaviorFrequency", string(data)}, " ")
}
