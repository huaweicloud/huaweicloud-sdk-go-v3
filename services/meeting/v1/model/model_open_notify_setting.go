package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// OpenNotifySetting 网络研讨会通知配置。
type OpenNotifySetting struct {

	// 发送邮件日历是否开启。 * Y：开启 * N： 不开启
	EnableCalendar string `json:"enableCalendar"`

	// 发送短信通知是否开启。 * Y：开启 * N： 不开启
	EnableSms string `json:"enableSms"`

	// 发送邮件是否开启。 * Y：开启 * N： 不开启
	EnableEmail string `json:"enableEmail"`
}

func (o OpenNotifySetting) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "OpenNotifySetting struct{}"
	}

	return strings.Join([]string{"OpenNotifySetting", string(data)}, " ")
}
