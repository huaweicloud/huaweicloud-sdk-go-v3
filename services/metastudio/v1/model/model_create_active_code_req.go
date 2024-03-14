package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateActiveCodeReq 创建激活码请求。
type CreateActiveCodeReq struct {

	// 应用ID。
	RobotId string `json:"robot_id"`

	// 智能交互对话ID。
	RoomId string `json:"room_id"`

	// 有效天数（0表示长期有效）。
	ValidPeriod int32 `json:"valid_period"`
}

func (o CreateActiveCodeReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateActiveCodeReq struct{}"
	}

	return strings.Join([]string{"CreateActiveCodeReq", string(data)}, " ")
}
