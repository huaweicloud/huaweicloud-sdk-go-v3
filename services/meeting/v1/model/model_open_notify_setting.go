package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 通知配置
type OpenNotifySetting struct {
	// 发送邮件日历是否开启，\"Y\" 开启，\"N\" 不开启。

	EnableCalendar string `json:"enableCalendar"`
	// 短信通知是否开启，\"Y\" 开启，\"N\" 不开启。

	EnableSms string `json:"enableSms"`
	// 短信通知是否开启，\"Y\" 开启，\"N\" 不开启。

	EnableEmail string `json:"enableEmail"`
}

func (o OpenNotifySetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenNotifySetting struct{}"
	}

	return strings.Join([]string{"OpenNotifySetting", string(data)}, " ")
}
