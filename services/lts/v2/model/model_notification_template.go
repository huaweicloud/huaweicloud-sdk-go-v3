package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type NotificationTemplate struct {
}

func (o NotificationTemplate) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationTemplate struct{}"
	}

	return strings.Join([]string{"NotificationTemplate", string(data)}, " ")
}
