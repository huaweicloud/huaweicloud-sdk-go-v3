package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RegisterVoiceResponseBodyResult 注册声音响应。
type RegisterVoiceResponseBodyResult struct {

	// 注册声音的名称。
	VoiceName *string `json:"voice_name,omitempty"`
}

func (o RegisterVoiceResponseBodyResult) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RegisterVoiceResponseBodyResult struct{}"
	}

	return strings.Join([]string{"RegisterVoiceResponseBodyResult", string(data)}, " ")
}
