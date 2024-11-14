package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SwitchSecondLevelMonitoringRequestBody struct {

	// 是否开启秒级监控。 - true: 开启; - false: 关闭。
	Enabled bool `json:"enabled"`
}

func (o SwitchSecondLevelMonitoringRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SwitchSecondLevelMonitoringRequestBody struct{}"
	}

	return strings.Join([]string{"SwitchSecondLevelMonitoringRequestBody", string(data)}, " ")
}
