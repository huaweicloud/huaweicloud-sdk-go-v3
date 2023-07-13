package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchSecondLevelMonitoringRequestBody struct {

	// 是否开启秒级监控。 取值为true为开启，取值为false为关闭。
	Enabled bool `json:"enabled"`
}

func (o SwitchSecondLevelMonitoringRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSecondLevelMonitoringRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchSecondLevelMonitoringRequestBody", string(data)}, " ")
}
