package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateWelcomeSpeechReq 创建欢迎词请求。
type CreateWelcomeSpeechReq struct {

	// 应用ID。
	RobotId string `json:"robot_id"`

	// 欢迎词。
	WelcomeSpeech string `json:"welcome_speech"`
}

func (o CreateWelcomeSpeechReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateWelcomeSpeechReq struct{}"
	}

	return strings.Join([]string{"CreateWelcomeSpeechReq", string(data)}, " ")
}
