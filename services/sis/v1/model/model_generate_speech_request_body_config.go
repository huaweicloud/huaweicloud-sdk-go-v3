package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GenerateSpeechRequestBodyConfig 语音合成配置信息。
type GenerateSpeechRequestBodyConfig struct {

	// 语音格式头：wav、mp3、pcm。 默认：wav
	AudioFormat *string `json:"audio_format,omitempty"`

	// 采样率：8kHz、16kHz、24kHz。 默认：24kHz
	SampleRate *string `json:"sample_rate,omitempty"`

	// 音色名称。
	VoiceName string `json:"voice_name"`

	// 语速：-500~500。 默认：0
	Speed *int32 `json:"speed,omitempty"`

	// 音高：-500~500。 默认：0
	Pitch *int32 `json:"pitch,omitempty"`

	// 音量：0~100。 默认：50
	Volume *int32 `json:"volume,omitempty"`
}

func (o GenerateSpeechRequestBodyConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GenerateSpeechRequestBodyConfig struct{}"
	}

	return strings.Join([]string{"GenerateSpeechRequestBodyConfig", string(data)}, " ")
}
