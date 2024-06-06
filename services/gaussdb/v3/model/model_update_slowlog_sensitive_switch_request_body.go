package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdateSlowlogSensitiveSwitchRequestBody struct {

	// 慢日志开关状态。
	OpenSlowLogSwitch bool `json:"open_slow_log_switch"`
}

func (o UpdateSlowlogSensitiveSwitchRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateSlowlogSensitiveSwitchRequestBody struct{}"
	}

	return strings.Join([]string{"UpdateSlowlogSensitiveSwitchRequestBody", string(data)}, " ")
}
