package model

import (
	"encoding/json"

	"strings"
)

type MultiAudio struct {
	// 音轨信息
	TracksInfo *[]TracksInfo `json:"tracks_info,omitempty"`
	// 音频文件
	AudioFiles *[]AudioFile `json:"audio_files,omitempty"`
	// 默认语言
	DefaultLanguage *string `json:"default_language,omitempty"`
}

func (o MultiAudio) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "MultiAudio struct{}"
	}

	return strings.Join([]string{"MultiAudio", string(data)}, " ")
}
