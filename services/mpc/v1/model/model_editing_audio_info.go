package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type EditingAudioInfo struct {

	// 音频编码格式
	Codec *string `json:"codec,omitempty"`

	// 音频采样率
	Sample *int32 `json:"sample,omitempty"`

	// 音频信道
	Channels *int32 `json:"channels,omitempty"`

	// 音频码率，单位: kbit/s
	Bitrate *int32 `json:"bitrate,omitempty"`

	// 音频码率，单位: bit/s
	BitrateBps *int64 `json:"bitrate_bps,omitempty"`
}

func (o EditingAudioInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EditingAudioInfo struct{}"
	}

	return strings.Join([]string{"EditingAudioInfo", string(data)}, " ")
}
