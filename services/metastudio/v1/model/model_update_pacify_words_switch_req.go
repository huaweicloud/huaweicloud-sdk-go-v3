package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePacifyWordsSwitchReq 修改安抚话术功能开关请求。
type UpdatePacifyWordsSwitchReq struct {

	// 应用ID。
	RobotId string `json:"robot_id"`

	Language *LanguageEnum `json:"language"`

	// 安抚话术功能开关。
	EnablePacifyWords bool `json:"enable_pacify_words"`
}

func (o UpdatePacifyWordsSwitchReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePacifyWordsSwitchReq struct{}"
	}

	return strings.Join([]string{"UpdatePacifyWordsSwitchReq", string(data)}, " ")
}
