package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationEndTimeResp **参数解释**： 告警通知关闭时间。    **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。
type NotificationEndTimeResp struct {
}

func (o NotificationEndTimeResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationEndTimeResp struct{}"
	}

	return strings.Join([]string{"NotificationEndTimeResp", string(data)}, " ")
}
