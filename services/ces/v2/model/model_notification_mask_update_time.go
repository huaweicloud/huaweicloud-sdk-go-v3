package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationMaskUpdateTime 告警屏蔽的更新时间，UNIX时间戳，单位毫秒。
type NotificationMaskUpdateTime struct {
}

func (o NotificationMaskUpdateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationMaskUpdateTime struct{}"
	}

	return strings.Join([]string{"NotificationMaskUpdateTime", string(data)}, " ")
}
