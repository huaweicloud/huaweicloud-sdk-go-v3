package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWelcomeSpeechReq 修改欢迎词请求。
type UpdateWelcomeSpeechReq struct {

	// 欢迎词。
	WelcomeSpeech *string `json:"welcome_speech,omitempty"`
}

func (o UpdateWelcomeSpeechReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWelcomeSpeechReq struct{}"
	}

	return strings.Join([]string{"UpdateWelcomeSpeechReq", string(data)}, " ")
}
