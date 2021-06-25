package model

import (
	"encoding/json"

	"strings"
)

type AudioTemplateInfo struct {
	// 音频采样率(有效值范围)<br/>  AUDIO_SAMPLE_AUTO=1 (default), AUDIO_SAMPLE_22050=2<br/> AUDIO_SAMPLE_32000=3<br/> AUDIO_SAMPLE_44100=4<br/> AUDIO_SAMPLE_48000=5<br/> AUDIO_SAMPLE_96000=6<br/>

	SampleRate int32 `json:"sample_rate"`
	// 音频码率（单位：Kbps）<br/>

	Bitrate *int32 `json:"bitrate,omitempty"`
	// 声道数(有效值范围)<br/> AUDIO_CHANNELS_1=1<br/> AUDIO_CHANNELS_2=2<br/>

	Channels int32 `json:"channels"`
}

func (o AudioTemplateInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "AudioTemplateInfo struct{}"
	}

	return strings.Join([]string{"AudioTemplateInfo", string(data)}, " ")
}
