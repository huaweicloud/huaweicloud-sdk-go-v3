package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type RtcHistoryUsage struct {

	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。格式为YYYY-MM-DD
	Date *string `json:"date,omitempty"`

	// 标清视频时长，单位秒
	VideoSd *int64 `json:"video_sd,omitempty"`

	// 高清视频时长，单位秒
	VideoHd *int64 `json:"video_hd,omitempty"`

	// 超高清视频时长，单位秒
	VideoHdp *int64 `json:"video_hdp,omitempty"`

	// 音频时长，单位秒
	Audio *int64 `json:"audio,omitempty"`

	// 音视频总时长，单位秒
	TotalDuration *int64 `json:"total_duration,omitempty"`
}

func (o RtcHistoryUsage) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RtcHistoryUsage struct{}"
	}

	return strings.Join([]string{"RtcHistoryUsage", string(data)}, " ")
}
