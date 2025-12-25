package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationMaskCreateTime **参数解释**： 告警屏蔽的创建时间，UNIX时间戳，单位毫秒。 **约束限制**： 不涉及。 **取值范围**: 不涉及。 **默认取值**： 不涉及。
type NotificationMaskCreateTime struct {
}

func (o NotificationMaskCreateTime) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationMaskCreateTime struct{}"
	}

	return strings.Join([]string{"NotificationMaskCreateTime", string(data)}, " ")
}
