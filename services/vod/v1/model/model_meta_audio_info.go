package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type MetaAudioInfo struct {

	// 音频编码格式
	Codec *string `json:"codec,omitempty"`

	// 音频采样率
	SamplingRate *int32 `json:"sampling_rate,omitempty"`

	// 音频码率，单位：bit/s
	Bitrate *int64 `json:"bitrate,omitempty"`
}

func (o MetaAudioInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "MetaAudioInfo struct{}"
	}

	return strings.Join([]string{"MetaAudioInfo", string(data)}, " ")
}
