package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type TriggerSetting struct {

	// 定时设置，cron格式描述 '* * * * * *'。cron 基于 UTC 时间。
	Cron *string `json:"cron,omitempty"`
}

func (o TriggerSetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "TriggerSetting struct{}"
	}

	return strings.Join([]string{"TriggerSetting", string(data)}, " ")
}
