package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type VoiceListResultVoices struct {

	// 注册成功的声音名称。
	VoiceName *string `json:"voice_name,omitempty"`

	// 注册成功的声音语种。
	Language *string `json:"language,omitempty"`
}

func (o VoiceListResultVoices) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VoiceListResultVoices struct{}"
	}

	return strings.Join([]string{"VoiceListResultVoices", string(data)}, " ")
}
