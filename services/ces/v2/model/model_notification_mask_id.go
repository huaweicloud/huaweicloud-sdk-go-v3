package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationMaskId **参数解释**： 屏蔽规则ID **约束限制**： 不涉及 **取值范围**： 以nm开头，后跟[0,62]位字母或数字。 **默认取值**： 不涉及
type NotificationMaskId struct {
}

func (o NotificationMaskId) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationMaskId struct{}"
	}

	return strings.Join([]string{"NotificationMaskId", string(data)}, " ")
}
