package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdatePacifyWordsTriggerTimeReq 修改安抚话术等待触发时长。
type UpdatePacifyWordsTriggerTimeReq struct {

	// 应用ID。
	RobotId string `json:"robot_id"`

	Language *LanguageEnum `json:"language"`

	// 安抚话术等待触发时长，单位毫秒
	TriggerTime int32 `json:"trigger_time"`
}

func (o UpdatePacifyWordsTriggerTimeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePacifyWordsTriggerTimeReq struct{}"
	}

	return strings.Join([]string{"UpdatePacifyWordsTriggerTimeReq", string(data)}, " ")
}
