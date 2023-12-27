package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ReplyAudioInfo 回复音频信息
type ReplyAudioInfo struct {

	// 音频URL
	AudioUrl *string `json:"audio_url,omitempty"`

	// 音频名
	AudioName *string `json:"audio_name,omitempty"`
}

func (o ReplyAudioInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ReplyAudioInfo struct{}"
	}

	return strings.Join([]string{"ReplyAudioInfo", string(data)}, " ")
}
