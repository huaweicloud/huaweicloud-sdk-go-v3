package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowWelcomeSpeechSwitchResponse Response Object
type ShowWelcomeSpeechSwitchResponse struct {

	// 欢迎词功能开关。
	EnableWelcomeSpeech *bool `json:"enable_welcome_speech,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowWelcomeSpeechSwitchResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowWelcomeSpeechSwitchResponse struct{}"
	}

	return strings.Join([]string{"ShowWelcomeSpeechSwitchResponse", string(data)}, " ")
}
