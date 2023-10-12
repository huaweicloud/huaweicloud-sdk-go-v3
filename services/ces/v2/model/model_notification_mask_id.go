package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationMaskId 屏蔽规则ID
type NotificationMaskId struct {
}

func (o NotificationMaskId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationMaskId struct{}"
	}

	return strings.Join([]string{"NotificationMaskId", string(data)}, " ")
}
