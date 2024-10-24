package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AudioSelectorSettings 设置音频选择器
type AudioSelectorSettings struct {
	AudioLanguageSelection *AudioSelectorLangSelection `json:"audio_language_selection,omitempty"`

	AudioPidSelection *AudioSelectorPidSelection `json:"audio_pid_selection,omitempty"`
}

func (o AudioSelectorSettings) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AudioSelectorSettings struct{}"
	}

	return strings.Join([]string{"AudioSelectorSettings", string(data)}, " ")
}
