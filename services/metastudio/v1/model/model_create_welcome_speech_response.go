package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWelcomeSpeechResponse Response Object
type CreateWelcomeSpeechResponse struct {

	// 欢迎词ID。
	WelcomeSpeechId *string `json:"welcome_speech_id,omitempty"`

	// 欢迎词。
	WelcomeSpeech *string `json:"welcome_speech,omitempty"`

	// 欢迎词功能开关。
	EnableWelcomeSpeech *bool `json:"enable_welcome_speech,omitempty"`

	Language *LanguageEnum `json:"language,omitempty"`

	// 应用ID。
	RobotId *string `json:"robot_id,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateWelcomeSpeechResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWelcomeSpeechResponse struct{}"
	}

	return strings.Join([]string{"CreateWelcomeSpeechResponse", string(data)}, " ")
}
