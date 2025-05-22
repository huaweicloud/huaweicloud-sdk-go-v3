package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// EncoderSettingsExpand 输出编码扩展
type EncoderSettingsExpand struct {

	// 音频输出配置的描述信息
	AudioDescriptions *[]EncoderSettingsExpandAudioDescriptions `json:"audio_descriptions,omitempty"`

	VideoDescriptions *VideoDescriptions `json:"video_descriptions,omitempty"`
}

func (o EncoderSettingsExpand) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "EncoderSettingsExpand struct{}"
	}

	return strings.Join([]string{"EncoderSettingsExpand", string(data)}, " ")
}
