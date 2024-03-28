package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// TtsConfig 语音合成配置 - 目前可用tts资源尚未确定，校验待定，测试可以先不测后端接口校验
type TtsConfig struct {

	// 用于设置音色
	Property string `json:"property"`

	// 用户设置音速
	Speed float32 `json:"speed"`

	// 用于设置音量
	Volume int32 `json:"volume"`

	// 段首停顿时间。 范围：0~60； 单位：秒 默认：0
	Delay float32 `json:"delay"`

	// 音高。 取值范围： -500~500 默认值：0
	Pitch *string `json:"pitch,omitempty"`

	// 语音格式头：wav、mp3、pcm 默认：wav
	AudioFormat *string `json:"audio_format,omitempty"`

	// 采样率：16000、8000 默认：8000
	SampleRate *string `json:"sample_rate,omitempty"`

	// tts来源： 0：huawei 1：ali 2：用户克隆声音 默认：0
	TtsSource *int32 `json:"tts_source,omitempty"`
}

func (o TtsConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TtsConfig struct{}"
	}

	return strings.Join([]string{"TtsConfig", string(data)}, " ")
}
