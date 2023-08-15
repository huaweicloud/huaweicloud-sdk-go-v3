package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpStreamDetail struct {

	// 采样时间。日期格式按照ISO8601表示法，并使用UTC时间。 格式为：YYYY-MM-DDThh:mm:ssZ。
	Time *string `json:"time,omitempty"`

	// 帧率，单位为fps
	Fps *int64 `json:"fps,omitempty"`

	// 码率，单位为Kbps
	Rate *int64 `json:"rate,omitempty"`

	// 时延 单位ms
	Delay *int64 `json:"delay,omitempty"`

	// 最近一次gop的时长 单位ms
	GopDuration *int64 `json:"gop_duration,omitempty"`

	// 视频DTS 单位ms
	LastVideoPts *int64 `json:"last_video_pts,omitempty"`

	// 音频DTS 单位ms
	LastAudioPts *int64 `json:"last_audio_pts,omitempty"`

	// 音视频DTS差值 单位ms
	LastVideoAudioPtsDiff *int64 `json:"last_video_audio_pts_diff,omitempty"`
}

func (o UpStreamDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpStreamDetail struct{}"
	}

	return strings.Join([]string{"UpStreamDetail", string(data)}, " ")
}
