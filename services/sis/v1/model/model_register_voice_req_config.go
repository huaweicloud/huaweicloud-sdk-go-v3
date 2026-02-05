package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterVoiceReqConfig 配置信息
type RegisterVoiceReqConfig struct {

	// 注册声音的名称，不能以数字、下划线开头。仅包含大、小写英文字母、数字和下划线，且长度不超过20个字符。一个project id下的声音名称不能重复。
	VoiceName string `json:"voice_name"`

	// data中语音数据的语种，取值chinese、english。 默认是chinese。
	Language *string `json:"language,omitempty"`
}

func (o RegisterVoiceReqConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterVoiceReqConfig struct{}"
	}

	return strings.Join([]string{"RegisterVoiceReqConfig", string(data)}, " ")
}
