package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationBeginTimeResp **参数解释**： 告警通知开启时间。如 00:00    **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。
type NotificationBeginTimeResp struct {
}

func (o NotificationBeginTimeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationBeginTimeResp struct{}"
	}

	return strings.Join([]string{"NotificationBeginTimeResp", string(data)}, " ")
}
