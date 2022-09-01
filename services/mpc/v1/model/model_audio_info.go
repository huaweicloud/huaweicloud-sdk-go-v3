package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AudioInfo struct {

	// 音频编码格式
	Codec *string `json:"codec,omitempty" xml:"codec"`

	// 音频采样率
	Sample *int32 `json:"sample,omitempty" xml:"sample"`

	// 音频信道
	Channels *int32 `json:"channels,omitempty" xml:"channels"`

	// 音频码率，单位: kbit/s
	Bitrate *int32 `json:"bitrate,omitempty" xml:"bitrate"`

	// 音频码率，单位: bit/s
	BitrateBps *int64 `json:"bitrate_bps,omitempty" xml:"bitrate_bps"`
}

func (o AudioInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioInfo struct{}"
	}

	return strings.Join([]string{"AudioInfo", string(data)}, " ")
}
