package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PutVideoInfo struct {
	TtsConfig *TtsConfig `json:"tts_config"`

	VideoConfig *VideoConfig `json:"video_config"`

	CharacterConfig *CharacterConfig `json:"character_config"`

	ReadConfig *ReadConfig `json:"read_config"`
}

func (o PutVideoInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PutVideoInfo struct{}"
	}

	return strings.Join([]string{"PutVideoInfo", string(data)}, " ")
}
