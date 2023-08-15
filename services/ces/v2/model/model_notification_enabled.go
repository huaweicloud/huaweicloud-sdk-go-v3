package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationEnabled 是否开启告警通知
type NotificationEnabled struct {
}

func (o NotificationEnabled) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationEnabled struct{}"
	}

	return strings.Join([]string{"NotificationEnabled", string(data)}, " ")
}
