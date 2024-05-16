package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateWelcomeSpeechSwitchReq 修改欢迎词功能开关请求。
type UpdateWelcomeSpeechSwitchReq struct {

	// 应用ID。
	RobotId string `json:"robot_id"`

	// 欢迎词功能开关。
	EnableWelcomeSpeech bool `json:"enable_welcome_speech"`
}

func (o UpdateWelcomeSpeechSwitchReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateWelcomeSpeechSwitchReq struct{}"
	}

	return strings.Join([]string{"UpdateWelcomeSpeechSwitchReq", string(data)}, " ")
}
