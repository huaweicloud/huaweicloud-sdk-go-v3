package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TrainingVideoMarkInfo 分身数字训练视频标记信息。
type TrainingVideoMarkInfo struct {

	// 训练视频起始时间。格式：“HH:MM:SS.mmm”。
	VideoStartTime *string `json:"video_start_time,omitempty"`

	// 训练视频结束时间。格式：“HH:MM:SS.mmm”。
	VideoEndTime *string `json:"video_end_time,omitempty"`
}

func (o TrainingVideoMarkInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TrainingVideoMarkInfo struct{}"
	}

	return strings.Join([]string{"TrainingVideoMarkInfo", string(data)}, " ")
}
