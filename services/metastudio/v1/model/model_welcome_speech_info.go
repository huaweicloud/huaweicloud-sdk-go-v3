package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// WelcomeSpeechInfo 欢迎词基本信息。
type WelcomeSpeechInfo struct {

	// 欢迎词ID。
	WelcomeSpeechId *string `json:"welcome_speech_id,omitempty"`

	// 欢迎词。
	WelcomeSpeech *string `json:"welcome_speech,omitempty"`

	// 应用ID。
	RobotId *string `json:"robot_id,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o WelcomeSpeechInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "WelcomeSpeechInfo struct{}"
	}

	return strings.Join([]string{"WelcomeSpeechInfo", string(data)}, " ")
}
