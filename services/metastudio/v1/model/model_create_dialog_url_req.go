package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateDialogUrlReq 创建对话链接。
type CreateDialogUrlReq struct {

	// 智能交互对话ID。
	RoomId string `json:"room_id"`

	// 应用ID。
	RobotId string `json:"robot_id"`
}

func (o CreateDialogUrlReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateDialogUrlReq struct{}"
	}

	return strings.Join([]string{"CreateDialogUrlReq", string(data)}, " ")
}
