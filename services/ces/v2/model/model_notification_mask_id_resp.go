package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationMaskIdResp **参数解释**： 屏蔽规则ID **取值范围**： 以nm开头，后跟[0,62]位字母或数字。
type NotificationMaskIdResp struct {
}

func (o NotificationMaskIdResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationMaskIdResp struct{}"
	}

	return strings.Join([]string{"NotificationMaskIdResp", string(data)}, " ")
}
