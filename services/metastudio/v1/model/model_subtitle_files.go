package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SubtitleFiles struct {
	TextSubtitleFile *SubtitleFileInfo `json:"text_subtitle_file,omitempty"`

	AudioSubtitleFile *SubtitleFileInfo `json:"audio_subtitle_file,omitempty"`
}

func (o SubtitleFiles) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SubtitleFiles struct{}"
	}

	return strings.Join([]string{"SubtitleFiles", string(data)}, " ")
}
