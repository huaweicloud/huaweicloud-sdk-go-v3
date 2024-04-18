package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AudioAssetMeta 音频元数据，自动提取获得。
type AudioAssetMeta struct {

	// 时长,单位秒
	Duration *int32 `json:"duration,omitempty"`

	// 音频编码格式
	AudioCodec *string `json:"audio_codec,omitempty"`

	// 音频平均码率,单位kbps
	AudioBitRate *int32 `json:"audio_bit_rate,omitempty"`

	// 音频声道数
	AudioChannels *int32 `json:"audio_channels,omitempty"`

	// 采样率,HZ
	Sample *int32 `json:"sample,omitempty"`

	ErrorInfo *ErrorResponse `json:"error_info,omitempty"`
}

func (o AudioAssetMeta) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioAssetMeta struct{}"
	}

	return strings.Join([]string{"AudioAssetMeta", string(data)}, " ")
}
