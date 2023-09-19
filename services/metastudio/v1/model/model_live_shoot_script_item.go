package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveShootScriptItem 直播话术配置。
type LiveShootScriptItem struct {

	// 剧本序号。
	SequenceNo *int32 `json:"sequence_no,omitempty"`

	// 段落标题。
	Title *string `json:"title,omitempty"`

	TextConfig *TextConfig `json:"text_config,omitempty"`

	AudioConfig *LiveAudioConfig `json:"audio_config,omitempty"`
}

func (o LiveShootScriptItem) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LiveShootScriptItem struct{}"
	}

	return strings.Join([]string{"LiveShootScriptItem", string(data)}, " ")
}
