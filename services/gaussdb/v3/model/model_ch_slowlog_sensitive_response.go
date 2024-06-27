package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ChSlowlogSensitiveResponse struct {

	// 慢日志脱敏开关。
	OpenSlowLogSwitch string `json:"open_slow_log_switch"`
}

func (o ChSlowlogSensitiveResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ChSlowlogSensitiveResponse struct{}"
	}

	return strings.Join([]string{"ChSlowlogSensitiveResponse", string(data)}, " ")
}
