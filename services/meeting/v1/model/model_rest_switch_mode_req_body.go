package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RestSwitchModeReqBody 切换会议显示策略请求。
type RestSwitchModeReqBody struct {

	// 会议显示策略。 - Fixed: 固定广播与会者 - VAS: 声控切换
	SwitchMode string `json:"switchMode"`

	// 画面类型。单画面设置只针对声控模式。 - 0: 单画面 - 1: 多画面
	ImageType int32 `json:"imageType"`
}

func (o RestSwitchModeReqBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RestSwitchModeReqBody struct{}"
	}

	return strings.Join([]string{"RestSwitchModeReqBody", string(data)}, " ")
}
