package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// NotificationEnabledResp **参数解释**： 是否开启告警通知。     **取值范围**： 布尔值。 - true:开启。 - false:关闭。
type NotificationEnabledResp struct {
}

func (o NotificationEnabledResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "NotificationEnabledResp struct{}"
	}

	return strings.Join([]string{"NotificationEnabledResp", string(data)}, " ")
}
