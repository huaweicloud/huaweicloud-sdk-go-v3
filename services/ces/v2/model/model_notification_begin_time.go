package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationBeginTime **参数解释**： 告警通知开启时间。    **约束限制**： 不涉及。 **取值范围**： 只能包含数字、“:”，长度为[1,64]个字符。           **默认取值**： 不涉及。
type NotificationBeginTime struct {
}

func (o NotificationBeginTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationBeginTime struct{}"
	}

	return strings.Join([]string{"NotificationBeginTime", string(data)}, " ")
}
