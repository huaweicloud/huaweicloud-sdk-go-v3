package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateHotWordsSwitchReq 修改热词功能开关请求。
type UpdateHotWordsSwitchReq struct {

	// 应用ID。
	RobotId string `json:"robot_id"`

	// 热词功能开关。
	EnableHotWords bool `json:"enable_hot_words"`
}

func (o UpdateHotWordsSwitchReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateHotWordsSwitchReq struct{}"
	}

	return strings.Join([]string{"UpdateHotWordsSwitchReq", string(data)}, " ")
}
