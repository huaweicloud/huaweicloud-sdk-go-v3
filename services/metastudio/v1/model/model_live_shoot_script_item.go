package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// LiveShootScriptItem 直播话术配置。
type LiveShootScriptItem struct {

	// **参数解释**： 剧本序号。 **约束限制**： 不涉及
	SequenceNo *int32 `json:"sequence_no,omitempty"`

	// **参数解释**： 段落标题。 **约束限制**： 不涉及 **取值范围**： 字符长度0-256位。 **默认取值**： 不涉及。
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
