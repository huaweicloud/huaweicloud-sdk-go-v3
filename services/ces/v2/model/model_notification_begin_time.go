package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationBeginTime 告警通知开启时间
type NotificationBeginTime struct {
}

func (o NotificationBeginTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationBeginTime struct{}"
	}

	return strings.Join([]string{"NotificationBeginTime", string(data)}, " ")
}
